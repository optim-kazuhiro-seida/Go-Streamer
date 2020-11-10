use crate::service::template::TEMPLATE;
use std::fs::File;
use std::io::{self, Write};
use std::path::PathBuf;

pub fn create_struct_list_service(str: &String) -> Vec<String> {
    return str
        .lines()
        .filter(|v| v.contains("struct"))
        .map(|t| {
            let strs: Vec<&str> = t.split_whitespace().collect();
            for i in 0..strs.len() - 1 {
                if i < strs.len() - 2 && strs[i + 1] == "struct" {
                    return String::from(strs[i]);
                }
            }
            return String::new();
        })
        .filter(|v| !v.is_empty())
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
