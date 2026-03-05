from pydantic_settings import BaseSettings

class Settings (BaseSettings):
    APP_NAME: str = "FastAPI Lev2"
    APP_VERSION: str = "0.1.0" 
    DEBUG: bool = True
    API_V1_PREFIX: str = "/api/v1"

    class config :
        env_file = ".env"
        env_file_encoding = "utf-8"

settings = Settings()