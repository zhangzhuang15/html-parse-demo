use std::borrow::Borrow;
use std::io;
use actix_web::{ App, HttpServer, Responder, Result };
use actix_web::web::Path;
use actix_files::NamedFile;


const addr: &str = "localhost:8077";

#[actix_web::get("/js/{file_name}")]
async fn get_js_file(file_name: Path<String>) -> Result<impl Responder> {
    let file_name = file_name.into_inner();
    let file = NamedFile::open_async(format!("./js/{}", file_name)).await?;
    Ok(file)
}

#[actix_web::get("/{html_name}")]
async fn get_html_file(html_name: Path<String>) -> Result<impl Responder> {
    let file_name = html_name.into_inner();
    let file = NamedFile::open_async(format!("./{}", file_name)).await?;
    Ok(file)
}

#[actix_web::main]
async fn main() -> io::Result<()> {
    let tip = format!("
server is running
please visit http://{}/index.html
", addr);

    println!("{}", tip);

    HttpServer::new(|| {
        App::new().service(get_js_file).service(get_html_file)
    })
    .bind(addr)?
    .workers(4)
    .run()
    .await
}