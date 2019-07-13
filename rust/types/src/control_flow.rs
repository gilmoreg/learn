#![allow(dead_code)]
#![allow(unused_variables)]

// If statement
fn weather(temp: i32) {
    if temp > 30 {
        println!("{} Hot outside", temp);
    } else if temp < 10 {
        println!("{} Cold outside", temp);
    } else {
        println!("{} Moderate outside", temp);
    }

    let day = if temp > 20 { "sunny" } else { "cloudy" };
    println!("{} day", day);

    // inline
    println!(
        "{} day",
        if temp > 20 {
            "hot"
        } else if temp < 10 {
            "cold"
        } else {
            "moderate"
        }
    );

    // nested inline
    println!(
        "{} day",
        if temp > 20 {
            if temp > 30 {
                "very hot"
            } else {
                "hot"
            }
        } else if temp < 10 {
            "cold"
        } else {
            "ok"
        }
    );
}

fn while_loop() {
    let mut x = 1;

    while x < 1000 {
        x *= 2;
        if x == 64 {
            continue;
        }
        println!("x {}", x);
    }

    let mut y = 1;
    // while true
    loop {
        y *= 2;
        println!("y {}", y);
        if y == 1 << 10 {
            break;
        }
    }

    // for loop different from other languages
    // stops at 10, not 11, double dots is exclusive (]
    for z in 1..11 {
        println!("z {}", z);
    }

    // also get index
    for (pos, a) in (30..41).enumerate() {
        println!("{}: {}", pos, a);
    }
}

fn match_statement() {
    let country_code = 44;

    let country = match country_code {
        44 => "UK",
        46 => "Sweden",
        7 => "Russia",
        1...999 => "Unknown", // triple dots, inclusive
        _ => "Invalid",
    };

    println!("country {}", country);
}

pub fn demo() {
    weather(5);
    weather(50);
    weather(35);
    while_loop();
    match_statement();
}