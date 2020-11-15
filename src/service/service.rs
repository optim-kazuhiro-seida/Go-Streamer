use crate::service::template::{TEMPLATE, TEMPLATE_FUNC};
use regex::Regex;
use std::fs::File;
use std::io::{self, Write};
use std::path::PathBuf;

pub fn create_struct_list_service(str: &String) -> Vec<String> {
    let i = Regex::new(r"[ ]*\S+[ ]+struct[ ]*\{").unwrap();
    return Regex::new(r"type[ ]*\S+[ ]+struct[ ]*\{")
        .unwrap()
        .find_iter(str.as_str())
        .map(|v| {
            v.as_str()
                .replace("struct", "")
                .replace("type", "")
                .replace("{", "")
                .replace(" ", "")
        })
        .chain(
            Regex::new(r"type[ \t]*\([\r\n \t\S]*?\)")
                .unwrap()
                .find_iter(str.as_str())
                .flat_map(|v| {
                    println!("{}", v.as_str());
                    return i.find_iter(v.as_str());
                })
                .map(|v| {
                    v.as_str()
                        .replace("struct", "")
                        .replace(" ", "")
                        .replace("{", "")
                }),
        )
        .collect();
}

pub fn save_struct_stream(struct_str: &String, package: &String, path: &String) {
    let mut p = PathBuf::from(path.as_str());
    p.set_file_name(format!(
        "{}{}{}",
        "stream_",
        struct_str.to_lowercase(),
        ".go"
    ));
    match File::create(p) {
        Ok(mut file) => {
            file.write_all(
                TEMPLATE
                    .to_string()
                    .replace("{{.TypeName}}", struct_str.as_str())
                    .replace("{{.PackageName}}", package)
                    .as_bytes(),
            )
            .expect("Fail save file.");
            println!("{} {}", "Completed ", struct_str.to_lowercase())
        }
        Err(err) => eprintln!("{}", err),
    }
}
pub fn save_struct_func(struct_str: &String, package: &String, path: &String) {
    let mut p = PathBuf::from(path.as_str());
    p.set_file_name(format!(
        "{}{}{}",
        "function_",
        struct_str.to_lowercase(),
        ".go"
    ));
    match File::create(p) {
        Ok(mut file) => {
            file.write_all(
                TEMPLATE_FUNC
                    .to_string()
                    .replace("{{.TypeName}}", struct_str.as_str())
                    .replace("{{.PackageName}}", package)
                    .as_bytes(),
            )
            .expect("Fail save file.");
            println!("{} {}", "Completed Func", struct_str.to_lowercase())
        }
        Err(err) => eprintln!("{}", err),
    }
}
