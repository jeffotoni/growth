import asyncio
import time

from fastapi import BackgroundTasks, FastAPI, Request, status
from fastapi.exceptions import HTTPException
from fastapi.routing import APIRouter
from pydantic import BaseModel

map_grow: dict[str, str] = {}
grow_lock = asyncio.Lock()


class DataGrowth(BaseModel):
    country: str
    indicator: str
    value: float
    year: int

    class Config:
        fields = {
            "country": "Country",
            "indicator": "Indicator",
            "value": "Value",
            "year": "Year",
        }


class StatusResponse(BaseModel):
    msg: str
    test_value: str
    count: int


class SizeResponse(BaseModel):
    value: int


class PutPayload(BaseModel):
    value: float


async def batch_store(grows: list[DataGrowth]):
    for data in grows:
        await store(data)


async def store(data: DataGrowth):
    key = f"{data.country.upper()}{data.indicator.upper()}{data.year}"
    async with grow_lock:
        map_grow[key] = data.value


app = FastAPI()
api_router = APIRouter(prefix="/api/v1")
growth_router = APIRouter(prefix="/growth", tags=["growth"])


@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    start_time = time.perf_counter()
    response = await call_next(request)
    process_time = time.perf_counter() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response


@app.get("/ping", tags=["healthcheck"])
async def ping():
    return "pong"


@growth_router.post("/", status_code=status.HTTP_202_ACCEPTED)
async def post(background_tasks: BackgroundTasks, grow: list[DataGrowth]):
    background_tasks.add_task(batch_store, grow)
    return {"msg": "In progress"}


@growth_router.get("/size", response_model=SizeResponse)
async def get_size():
    async with grow_lock:
        return SizeResponse(value=len(map_grow))


@growth_router.get("/status", response_model=StatusResponse)
async def get_status():
    key = "BRZNGDP_R2002"
    try:
        async with grow_lock:
            value = map_grow[key]
            return StatusResponse(
                msg="complete",
                test_value=f"{value:.2f}",
                count=len(map_grow),
            )
    except KeyError:
        return StatusResponse(
            msg="not finished",
            test_value="0.0",
            count=len(map_grow),
        )


@growth_router.delete(
    "/{country}/{indicator}/{year}",
    status_code=status.HTTP_204_NO_CONTENT,
)
async def delete(country: str, indicator: str, year: int):
    country = country.upper()
    indicator = indicator.upper()
    key = f"{country}{indicator}{year}"
    try:
        async with grow_lock:
            del map_grow[key]
    except KeyError:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Code not found"
        )


@growth_router.get("/{country}/{indicator}/{year}", response_model=DataGrowth)
async def get(country: str, indicator: str, year: int):
    country = country.upper()
    indicator = indicator.upper()
    key = f"{country}{indicator}{year}"
    try:
        async with grow_lock:
            value = map_grow[key]
            return DataGrowth(
                Country=country, Indicator=indicator, Value=value, Year=year
            )
    except KeyError:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND, detail="Code not found"
        )


@growth_router.put("/{country}/{indicator}/{year}", response_model=DataGrowth)
async def put(country: str, indicator: str, year: int, payload: PutPayload):
    data = DataGrowth(
        Country=country, Indicator=indicator, Value=payload.value, Year=year
    )
    await store(data)
    return data


api_router.include_router(growth_router)
app.include_router(api_router)
app.include_router(growth_router)
