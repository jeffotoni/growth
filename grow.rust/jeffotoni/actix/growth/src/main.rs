use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
//use futures::StreamExt;
use serde::{Deserialize, Serialize};

const MAX_SIZE: usize = 10485760_144;

#[derive(Serialize, Deserialize)]
struct Growth {
    Country: String,
    Indicator: String,
    Value: f32,
    Year: i8
}

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok()
    .content_type("application/json")
    .body("Hello world!")
}

//#[post("/api/v1/growth")]
// async fn post(info: web::Json<Growth>) -> Result<String, Error> {
//     Ok(format!("Welcome {}!", info.Country))
// }
async fn post(grow: web::Json<Growth>) -> impl Responder {
    //HttpResponse::Ok().body(grow)
    format!("Welcome {}!", grow.Country)
}

async fn ping() -> impl Responder {
    HttpResponse::Ok()
    .content_type("application/json")
    .body("{\"msg\":\"pong\"}")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .app_data(web::JsonConfig::default().limit(10485760))
            .service(hello)
            //.service(post)
            .route("/api/v1/growth", web::post().to(post))
            .route("/ping", web::get().to(ping))
    })
    .bind("0.0.0.0:8080")?
    .run()
    .await
}