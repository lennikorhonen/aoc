use std::collections::HashMap;

use read_file::file_to_str;

fn main() {
    let input = file_to_str(String::from("/home/lenni/coding/aoc/2023/day_2_rust/input"));
    // let test_input = file_to_str(String::from("/home/lenni/coding/aoc/2023/day_2_rust/test1.txt"));

    // let result = part_one(input);

    // let sum: u32 = result.iter().sum();
    // println!("Part one: {sum}");

    let result_two = part_two(input);

    let sum_two: u32 = result_two.iter().sum();
    println!("Part two: {sum_two}");
}

fn part_one(input: Vec<String>) -> Vec<u32> {
    let mut results = Vec::new();
    let start_value = HashMap::from([
        ("red", 12),
        ("green", 13),
        ("blue", 14)
    ]);

    for line in input {
        let mut line_vec = line.split(": ").collect::<Vec<&str>>();

        line_vec.remove(0);
        let game_sets_valid = validate_sets(line_vec[0], &start_value);

        if game_sets_valid {
            let game_id = get_game_id(line);

            results.push(game_id.parse::<u32>().unwrap());
        };
    }

    results
}

fn part_two(input: Vec<String>) -> Vec<u32> {
    let mut results = Vec::new();

    for line in input {
        let mut line_vec = line.split(": ").collect::<Vec<&str>>();
        
        line_vec.remove(0);

        let cube_power = loop_sets(line_vec[0]);

        results.push(cube_power);
    }

    results
}

fn get_game_id(line: String) -> String {
    line.clone()
        .find("Game ")
        .map(|index| line[index + "Game ".len()..index + "Game ".len() + 2]
            .to_string())
        .unwrap_or_default()
        .split(":")
        .collect::<Vec<_>>()
        .join("")
        .trim()
        .to_string()
}

fn validate_sets(sets: &str, start_values: &HashMap<&str, u32>) -> bool {
    for set in sets.split("; ") {
        if !check_set(set.to_string(), &start_values) {
            return false
        }
    }

    true
}

fn check_set(set: String, start_values: &HashMap<&str, u32>) -> bool {
    for set_values in set.split(", ") {
        let set_value = set_values.split(" ").collect::<Vec<&str>>();
        if Some(&set_value[0].parse::<u32>().unwrap()) > start_values.get(&set_value[1]) {
            return false
        }
    }
    true
}

fn loop_sets(sets: &str) -> u32 {
    let mut cubes = HashMap::from([
        (String::from("red"), 0),
        (String::from("green"), 0),
        (String::from("blue"), 0)
    ]);

    for set in sets.split("; ") {
        cubes = least_num_of_cubes(set.to_string(), cubes);
    }

    let cubes_vec: Vec<u32> = cubes
        .into_values()
        .collect();

    cubes_vec[0] * cubes_vec[1] * cubes_vec[2]
}

fn least_num_of_cubes(set: String, mut least_cubes: HashMap<String, u32>) -> HashMap<String, u32> {

    for set_values in set.split(", ") {
        least_cubes = update_map(set_values, least_cubes);
    }

    least_cubes
}

fn update_map(set_values: &str, mut map: HashMap<String, u32>) -> HashMap<String, u32> {
    let set_value: Vec<&str> = set_values.split(" ").collect();

    if map.get(set_value[1]).unwrap() < &set_value[0].parse::<u32>().unwrap() {
        map.insert(set_value[1].to_string(), set_value[0].parse::<u32>().unwrap());
    }

    map
}
