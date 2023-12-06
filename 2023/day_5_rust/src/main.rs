use std::collections::HashMap;

use read_file::file_to_str;

fn main() {
    let input = file_to_str(String::from("test5.txt"));

    part_one(input);
}

fn part_one(mut input: Vec<String>) -> u32 {
    let mut seeds_vec = input[0]
        .split(": ")
        .collect::<Vec<&str>>();

    seeds_vec.remove(0);

    let seeds_str = seeds_vec
        .join(" ")
        .to_string();

    let seeds: Vec<&str> = seeds_str
        .split(" ")
        .collect(); 

    println!("{:?}", seeds);

    let mut maps: Vec<Vec<String>> = Vec::new();
    let mut single_map = Vec::new();

    input.remove(0);
    input.remove(0);

    for line in input {
        println!("{line}");
        if line.is_empty() {
            maps.push(single_map);
            single_map = Vec::new();
            continue;
        }

        single_map.push(line);
    }

    src_to_dest(maps);
    1
}

fn src_to_dest(maps: Vec<Vec<String>>) {
    let mut source_to_destination: HashMap<&str, HashMap<&str, &str>> = HashMap::new();
    let mut map_name = String::from("");

    println!("{:?}", maps);

    for map in maps {
        let vals: HashMap<&str, &str> = HashMap::new();

        println!("{:?}", map);

        map_name = map[0];

        for val in map {
            if val.contains(":") {
                continue;
            }

            // let val_vec: Vec<&str> = val.split(" ").collect();

            // vals.insert("dest", val_vec[0]);
            // vals.insert("src", val_vec[1]);
            // vals.insert("len", val_vec[2]);

        }

        source_to_destination.insert(&map_name, vals);
    }
}

