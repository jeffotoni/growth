use actix_web::{
    error, middleware, web, App, Error, HttpRequest, HttpResponse, HttpServer,
};
use futures::StreamExt;
use json::JsonValue;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
struct Growth {
    country: String,
    indicator: String,
    value: f32,
    year: i8,
}

/// This handler uses json extractor
async fn index(item: web::Json<Growth>) -> HttpResponse {
    println!("model: {:?}", &item);
    HttpResponse::Ok().json(item.0) // <- send response
}

/// This handler uses json extractor with limit
async fn extract_item(item: web::Json<Growth>, req: HttpRequest) -> HttpResponse {
    println!("request: {:?}", req);
    println!("model: {:?}", item);

    HttpResponse::Ok().json(item.0) // <- send json response
}

const MAX_SIZE: usize = 10485760_144; // max payload size is 256k

/// This handler manually load request payload and parse json object
async fn index_manual(mut payload: web::Payload) -> Result<HttpResponse, Error> {
    // payload is a stream of Bytes objects
    let mut body = web::BytesMut::new();
    while let Some(chunk) = payload.next().await {
        let chunk = chunk?;
        // limit max size of in-memory payload
        if (body.len() + chunk.len()) > MAX_SIZE {
            return Err(error::ErrorBadRequest("overflow"));
        }
        body.extend_from_slice(&chunk);
    }

    // body is loaded, now we can deserialize serde-json
    let obj = serde_json::from_slice::<Growth>(&body)?;
    Ok(HttpResponse::Ok().json(obj)) // <- send response
}

/// This handler manually load request payload and parse json-rust
async fn index_mjsonrust(body: web::Bytes) -> Result<HttpResponse, Error> {
    // body is loaded, now we can deserialize json-rust
    let result = json::parse(std::str::from_utf8(&body).unwrap()); // return Result
    let injson: JsonValue = match result {
        Ok(v) => v,
        Err(e) => json::object! {"err" => e.to_string() },
    };
    Ok(HttpResponse::Ok()
        .content_type("application/json")
        .body(injson.dump()))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_LOG", "actix_web=info");
    env_logger::init();

    HttpServer::new(|| {
        App::new()
            // enable logger
            .wrap(middleware::Logger::default())
            .data(web::JsonConfig::default().limit(10485760)) // <- limit size of the payload (global configuration)
            .service(web::resource("/extractor").route(web::post().to(index)))
            .service(
                web::resource("/extractor2")
                    .data(web::JsonConfig::default().limit(10485760)) // <- limit size of the payload (resource level)
                    .route(web::post().to(extract_item)),
            )
            .service(web::resource("/manual").route(web::post().to(index_manual)))
            .service(
                web::resource("/mjsonrust")
                    .data(web::JsonConfig::default().limit(10485760_144))
                    .route(web::post().to(index_mjsonrust))
                )
            .service(
                web::resource("/")
                    .data(web::JsonConfig::default().limit(10485760))
                    .route(web::post().to(index))
                )
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}

#[cfg(test)]
mod tests {
    use super::*;
    use actix_web::dev::Service;
    use actix_web::{http, test, web, App};

    #[actix_rt::test]
    async fn test_index() -> Result<(), Error> {
        let mut app = test::init_service(
            App::new().service(web::resource("/").route(web::post().to(index))),
        )
        .await;

        let req = test::TestRequest::post()
            .uri("/")
            .set_json(&Growth {
                name: "my-name".to_owned(),
                number: 43,
            })
            .to_request();
        let resp = app.call(req).await.unwrap();

        assert_eq!(resp.status(), http::StatusCode::OK);

        let response_body = match resp.response().body().as_ref() {
            Some(actix_web::body::Body::Bytes(bytes)) => bytes,
            _ => panic!("Response error"),
        };

        assert_eq!(response_body, r##"{"name":"my-name","number":43}"##);

        Ok(())
    }
}

// use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
// //use futures::StreamExt;
// use serde::{Deserialize, Serialize};

// const MAX_SIZE: usize = 10485760_144;

// #[derive(Serialize, Deserialize)]
// struct Growth {
//     Country: String,
//     Indicator: String,
//     Value: f32,
//     Year: i8
// }

// #[get("/")]
// async fn hello() -> impl Responder {
//     HttpResponse::Ok()
//     .content_type("application/json")
//     .body("Hello world!")
// }

// //#[post("/api/v1/growth")]
// // async fn post(info: web::Json<Growth>) -> Result<String, Error> {
// //     Ok(format!("Welcome {}!", info.Country))
// // }
// async fn post(grow: web::Json<Growth>) -> impl Responder {
//     //HttpResponse::Ok().body(grow)
//     format!("Welcome {}!", grow.Country)
// }

// async fn ping() -> impl Responder {
//     HttpResponse::Ok()
//     .content_type("application/json")
//     .body("{\"msg\":\"pong\"}")
// }

// #[actix_web::main]
// async fn main() -> std::io::Result<()> {
//     HttpServer::new(|| {
//         App::new()
//             .app_data(web::JsonConfig::default().limit(10485760))
//             .service(hello)
//             //.service(post)
//             .route("/api/v1/growth", web::post().to(post))
//             .route("/ping", web::get().to(ping))
//     })
//     .bind("0.0.0.0:8080")?
//     .run()
//     .await
// }