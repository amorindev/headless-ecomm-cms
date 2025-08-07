# Headless CMS for Mobile E-commerce

This is a headless content management system designed specifically for mobile-focused e-commerce platforms. It provides a flexible and modular backend to manage products, categories, inventory, and digital assets without being tied to any specific frontend technology.

## Description

The Headless E-commerce CMS enables developers and businesses to manage their store's content and inventory efficiently through a modern API-first architecture. Tailored for mobile applications, it supports seamless integration with mobile frontends (Flutter, React Native, etc.).

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/amorindev/headless-ecomm-cms.git
    ```

2. Download dependencies:
    ```bash
    go mod tidy
    ```

3. Set environment variables, add a `.env` file based on `env.example`:
4. Start development containers 
    ```bash
    make compose-dev
    ```
5. Expose MinIO via Cloudflare Tunnel
    ```bash
    make cloudflared-minio
    ```
6. Run the project:
    ```bash
    make run
    ```
