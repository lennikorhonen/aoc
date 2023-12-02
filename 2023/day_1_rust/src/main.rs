use read_file::file_to_str;

fn main() {
    let input1 = file_to_str(String::from("/home/lenni/coding/aoc/2023/day_1_rust/input1"));
    println!("{}", part_one(input1));
    
    let input2 = file_to_str(String::from("/home/lenni/coding/aoc/2023/day_1_rust/test2.txt"));
    println!("{}", part_two(input2));
}

fn part_one(input: Vec<String>) -> i32 {
    let mut values = Vec::new();

    for line in input {
        let number_str: String = vec![first_number(&line), last_number(&line)]
            .iter()
            .collect();

        let number: i32 = number_str.parse().unwrap();
        values.push(number);
    }
    // println!("{:?}", values);
    values.iter().sum()
}

fn part_two(input: Vec<String>) -> i32 {
    let mut values = Vec::new();

    for line in input {
        let number_str: String = vec![first_number(&line), last_number(&line)]
            .iter()
            .collect();

        let number: i32 = number_str.parse().unwrap();
        values.push(number);
    }
    values.iter().sum()
}

// Return first found i32
fn first_number(line: &str) -> char {
    let mut num_as_string_vec = Vec::new();

    for c in line.chars() {
        let num_as_string = &num_as_string_vec.iter().clone().collect::<String>();
        if strings_to_numbers(Some(&num_as_string)) != '0' {
            return strings_to_numbers(Some(&num_as_string))
        }

        let found = check_if_contains(&num_as_string);

        if strings_to_numbers(found) != '0' {
            return strings_to_numbers(found)
        }

        if c.is_digit(10) {
            return c
        } 

        num_as_string_vec.push(c);
    }

    '0'
}

// Return the last i32
fn last_number(line: &str) -> char {
    // ## Part One ##
    // first_number(&line.chars().rev().collect::<String>())

    // ## Part Two ##
    let mut num_as_string_vec = Vec::new();
    
    for c in line.chars().rev() {
        let num_as_string = reverse_num_str(&num_as_string_vec.iter().clone().collect::<String>());
        if strings_to_numbers(Some(&num_as_string)) != '0' {
            return strings_to_numbers(Some(&num_as_string))
        }

        let found = check_if_contains(&num_as_string);

        if strings_to_numbers(found) != '0' {
            return strings_to_numbers(found)
        }

        if c.is_digit(10) {
            return c
        }

        num_as_string_vec.push(c);
    }

    '0'
}

fn reverse_num_str(line: &str) -> String {
    line.chars().rev().collect::<String>()
}

fn strings_to_numbers(num_string: Option<&str>) -> char {
    match num_string {
        Some("one") => '1',
        Some("two") => '2',
        Some("three") => '3',
        Some("four") => '4',
        Some("five") => '5',
        Some("six") => '6',
        Some("seven") => '7',
        Some("eight") => '8',
        Some("nine") => '9',
        Some(&_) => '0',
        None => '0'
    }
}

fn check_if_contains(num_string: &str) -> Option<&str> {
    let num_vec = vec!["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];
    let mut result;

    for num_str in num_vec.into_iter() {
        result = num_string
            .find(num_str)
            .map(|index| &num_string[index..index + num_str.len()]);

        if result != None {
            return result
        }
    }

    None
}
