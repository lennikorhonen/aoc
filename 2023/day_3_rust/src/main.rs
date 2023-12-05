use read_file::file_to_str;

struct PartNum {
    num: i32,
    coords: Vec<Coords>
}

struct Coords {
    x: i32,
    y: i32
}

fn main() {
    let input = file_to_str(String::from("test3.txt"));

    println!("{:?}", input);
    part_one(input);
}

fn part_one(input: Vec<String>) -> Vec<u32> {
    let found: Vec<PartNum> = Vec::new();

    let mut x = 0;
    let mut y = 0;

    for line in input {
        x = 0;
        for c in line.chars() {
            println!("{x}, {y}, {c}");
            x += 1;
        }
        y += 1;
    }
    vec![1,2]
}

fn check_diagonal(curr_line: String, next_line: String) {}
fn check_adjacent_vertical(curr_line: String, next_line: String) {}
fn check_adjacent_horizontal(line: String) {
    let mut i = 1;

    let test: Vec<char> = line.chars().collect();

    for c in &test {
        if i < line.len() {
            print!("{c}, {}\n", &test[i]);
            i += 1;
        }
    }
}
