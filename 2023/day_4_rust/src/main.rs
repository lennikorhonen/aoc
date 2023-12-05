use std::collections::HashMap;

use read_file::file_to_str;

fn main() {
    let input1 = file_to_str(String::from("input"));
    let input = file_to_str(String::from("input"));

    part_one(input1);
    part_two(input);
}

fn part_one(input: Vec<String>) {
    let mut points = 0;

    for line in input {
        let mut card: Vec<&str> = line.split(": ").collect();

        card.remove(0);
        let card_num_str = card.join("");

        let card_nums: Vec<&str> = card_num_str.split(" | ").collect();

        let card_left = card_nums[0];
        let card_right = card_nums[1];

        let card_left_vec: Vec<&str> = card_left.split(" ").collect();
        let card_right_vec: Vec<&str> = card_right.split(" ").collect();

        let matches = left_num_in_right(card_left_vec, card_right_vec);

        points += calculate_points(matches);
    }
    println!("Points: {points}");
}

fn part_two(input: Vec<String>) {
    let mut card_instances: HashMap<u32, u32> = HashMap::new();

    let mut tot_rows = 1u32;
    while tot_rows <= input.len().try_into().unwrap() {
        card_instances.insert(tot_rows, 1);
        tot_rows += 1;
    }

    for line in input {
        let card: Vec<&str> = line.split(": ").collect();

        let card_num = card[0];
        let card_vals = card[1];

        let card_id = get_card_id(&card_num);

        let card_nums: Vec<&str> = card_vals.split(" | ").collect();

        let card_left = card_nums[0];
        let card_right = card_nums[1];

        let card_left_vec: Vec<&str> = card_left.split(" ").collect();
        let card_right_vec: Vec<&str> = card_right.split(" ").collect();

        let matches = left_num_in_right(card_left_vec, card_right_vec);

        card_instances = loop_key_val_pairs(matches, card_id, card_instances);
    }

    println!("{:?}", card_instances);

    let total_cards = calculate_total_cards(card_instances);
    println!("Total cards: {total_cards}");
}

fn left_num_in_right(card_left: Vec<&str>, card_right: Vec<&str>) -> u32 {
    let mut matches: u32 = 0;

    for left_num in card_left {
        if card_right.iter().any(|&i| i.trim() == left_num.trim()) {
            if left_num != "" {
                matches += 1;
            }
        }
    }
    matches
}

fn calculate_points(matches: u32) -> u32 {
    let mut points = 0;

    if matches != 0 {
        points += u32::pow(2, matches - 1);
    }

    points
}

fn calculate_total_cards(card_map: HashMap<u32, u32>) -> u32 {
    let total_cards = card_map
        .into_values()
        .into_iter()
        .sum();

    total_cards
}

fn get_card_id(card_num: &str) -> u32 {
    let mut card_vec: Vec<&str> = card_num.split(" ").collect();

    card_vec.remove(0);

    let card_id: u32 = card_vec.join("").parse().unwrap();
    card_id
}

fn loop_key_val_pairs(matches: u32, mut card_id: u32, card_instances: HashMap<u32, u32>) -> HashMap<u32, u32> {
    let mut i = 0;

    let mut temp_map = HashMap::from(card_instances.clone());
    let curr_card_amount = card_instances.get(&card_id).unwrap();

    while i < matches {

        let next_card_amount = card_instances.get(&(card_id+1)).unwrap();
        let new_amount = next_card_amount + curr_card_amount;

        temp_map = update_hashmap(temp_map, card_id, new_amount);

        card_id += 1;
        i += 1;
    }

    temp_map
}

fn update_hashmap(mut map: HashMap<u32, u32>, key: u32, val: u32) -> HashMap<u32, u32> {
    map.insert(key+1, val);

    map
}
