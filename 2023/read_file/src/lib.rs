use std::fs::read_to_string;

// Function to read in the AOC prompt file
pub fn file_to_str(file_path: String) -> Vec<String> {
    read_to_string(file_path)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

