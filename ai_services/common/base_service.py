from fastapi import FastAPI


def create_base_app(title: str) -> FastAPI:
    app = FastAPI(
        title=title,
        version="1.0",
        docs_url="/docs",
        redoc_url="/redoc",
    )
    return app
