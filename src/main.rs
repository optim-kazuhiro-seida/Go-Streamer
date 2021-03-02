mod service;
mod util;

use itertools::Itertools;
use service::service::{create_struct_list_service, save_struct_stream};
use std::env;
use util::file::read_file_infos;

fn main() {
    let dir = match env::var("STREAMER_DIRECTORY") {
        Ok(val) => val,
        _ => String::from("./"),
    };
    let recursion = match env::var("STREAMER_RECURSION") {
        Ok(val) => val == "",
        _ => false,
    };
    println!(
        "STREAMER_DIRECTORY: {}\nSTREAMER_RECURSION: {}",
        &dir, recursion
    );
    read_file_infos(&dir, recursion)
        .into_iter()
        .filter(|v| v.path.ends_with(".go"))
        .for_each(|f| {
            if let Some(package_line) = f.value.split("\n").find(|v| v.contains("package")) {
                if let Some(package) = package_line.split(" ").find(|v| !v.contains("package")) {
                    create_struct_list_service(&f.value)
                        .into_iter()
                        .unique()
                        .filter(|str| {
                            str.chars()
                                .all(|c| c.is_ascii_alphabetic() || c.is_ascii_alphanumeric())
                        })
                        .filter(|v| v.len() > 0 && v.chars().nth(0).unwrap().is_uppercase())
                        .for_each(|v| save_struct_stream(&v, &package.to_string(), &f.path));
                }
            }
        });
}
