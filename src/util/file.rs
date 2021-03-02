use std::fs;
use std::fs::{metadata, read_to_string};
use std::path::PathBuf;
pub struct FileStringInfo {
    pub path: String,
    pub value: String,
    pub child_dir: Vec<String>,
}
impl FileStringInfo {
    fn clone(&self) -> FileStringInfo {
        FileStringInfo {
            child_dir: self.child_dir.clone(),
            value: self.value.clone(),
            path: self.path.clone(),
        }
    }
}

pub fn read_file_infos(path: &String, recursion: bool) -> Vec<FileStringInfo> {
    let mut result: Vec<FileStringInfo> = vec![];
    let dir = read_file_info(&path);
    result.push(dir.clone());
    dir.child_dir.iter().for_each(|c| {
        if recursion || metadata(&c).unwrap().is_file() {
            println!("{}", &c);
            result.append(&mut read_file_infos(&c, recursion))
        }
    });
    return result;
}

pub fn read_file_info(path: &String) -> FileStringInfo {
    let text: String = match read_to_string(path) {
        Ok(content) => content,
        Err(_) => String::new(),
    };
    let abs_path: String = match fs::canonicalize(&PathBuf::from(path)) {
        Ok(v) => String::from(v.into_os_string().into_string().unwrap()),
        Err(_) => String::new(),
    };
    return FileStringInfo {
        path: String::from(abs_path),
        value: String::from(text),
        child_dir: get_dir_list(path),
    };
}

fn get_dir_list(path: &String) -> Vec<String> {
    let values = fs::read_dir(path);
    return if values.is_err() {
        vec![]
    } else {
        values
            .unwrap()
            .map(|v| v.unwrap().path().display().to_string())
            .collect()
    };
}
