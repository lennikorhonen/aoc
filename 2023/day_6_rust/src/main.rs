fn main() {
    // Didn't bother to read in and parse file as this was a small input
    // Test
    // let time_input = vec![7, 15, 30];
    // let dist_input = vec![9, 40, 200];

    // Test 2
    let time_input = vec![71530];
    let dist_input = vec![940200];

    // Input hidden as asked by AoC creators
    // Part 1
    // let time_input = vec![];
    // let dist_input = vec![];

    // Part 2
    // let time_input = vec![];
    // let dist_input = vec![];

    let mut ways_to_win: Vec<usize> = Vec::new();
    let mut dist_i: usize = 0;

    for time in time_input {
        let mut i: usize = 0;

        let mut num_ways_to_win: usize = 0;

        while i < time {
            if (time - (time - i)) * (time - i) > dist_input[dist_i] {
                num_ways_to_win += 1;
            }

            i  += 1;
        }

        ways_to_win.push(num_ways_to_win);

        dist_i += 1;
    }

    let mut result = 1;

    for val in ways_to_win.iter() {
        result = result * val;
    }
    
    println!("res: {result}");

}

