#![allow(dead_code)]
#![allow(unused_variables)]
use std::mem;

struct Point {
    x: f64,
    y: f64,
}

struct Line {
    start: Point,
    end: Point,
}

enum Color {
    Red,
    Green,
    Blue,
    RgbColor(u8, u8, u8), // tuple is unnamed
    Cmyk {
        cyan: u8,
        magenta: u8,
        yellow: u8,
        black: u8,
    }, // struct is named
}

fn structures() {
    let p = Point { x: 3.0, y: 4.0 };
    println!("p: x {} y {}", p.x, p.y);

    let p2 = Point { x: 5.0, y: 10.0 };
    let line = Line { start: p, end: p2 };
}

fn enums() {
    // let c: Color = Color::RgbColor(255, 255, 0);
    let c: Color = Color::Cmyk {
        cyan: 0,
        magenta: 0,
        yellow: 0,
        black: 255,
    };

    match c {
        Color::Red => println!("R"),
        Color::Green => println!("G"),
        Color::Blue => println!("B"),
        Color::RgbColor(0, 0, 0)
        | Color::Cmyk {
            cyan: _,
            magenta: _,
            yellow: _,
            black: 255,
        } => println!("black"),
        Color::RgbColor(r, g, b) => println!("{},{},{}", r, g, b),
        _ => (), // do nothing catchall
    }
}

fn option(y: f64) {
    // presence or absence of a particular value
    // Option<T>
    let x = 3.0;
    let result: Option<f64> = if y != 0.0 { Some(x / y) } else { None };
    println!("{:?}", result);

    // getting it out
    match result {
        Some(z) => println!("{}", z),
        None => println!("none"),
    }

    // or if/while let
    // if condition fails, treated as false
    if let Some(z) = result {
        println!("z = {}", z)
    } else {
        println!("z = none")
    };
}

fn arrays() {
    // arrays require knowing size in advance
    let mut a: [i32; 5] = [1, 2, 3, 4, 5]; // type is optional
    println!("a has {} elements, first element is {}", a.len(), a[0]);
    a[0] = 300;
    println!("a has {} elements, first element is {}", a.len(), a[0]);
    println!("{:?}", a);
    if a != [1, 2, 3, 4, 5] {
        println!("something changed");
    }

    let b = [1u16; 10];
    for i in 0..b.len() {
        println!("b {}", b[i]);
    }

    println!("b took up {} bytes", mem::size_of_val(&b)); // 20

    // 2d arrays or matrices
    let mtx: [[f32; 3]; 2] = [[1.0, 0.0, 0.0], [1.0, 0.0, 0.0]];
    println!("{:?}", mtx);
}

fn vectors() {
    let mut a = Vec::new();
    a.push(1);
    a.push(2);
    a.push(3);
    println!("a = {:?}", a);

    // indexes have to be usize
    let i: usize = 0;
    println!("a[0] = {}", a[i]);

    // get to avoid out of range panics
    match a.get(1) {
        Some(x) => println!("a[1] = {}", x),
        None => println!("a[1] = none"),
    }

    // pop also returns an option
    match a.pop() {
        Some(x) => println!("pop = {}", x),
        None => println!("pop = none"),
    }

    // iterating
    for x in &a {
        println!("{}", x);
    }

    // or with pop (will print in reverse order)
    while let Some(x) = a.pop() {
        println!("{}", x);
    }
}

fn use_slice(slice: &mut [i32]) {
    println!("{}, {}", slice[0], slice.len());
    slice[0] = 1234;
}

fn slices() {
    let mut data = [1, 2, 3, 4, 5];
    use_slice(&mut data[1..4]);
    println!("{:?}", data);
}

fn strings() {
    // strings are vectors of utf-8 characters
    let s: &'static str = "hello there"; // statically allocated
                                         // s = "abc" doesn't work
                                         // indexing doesn't work because utf-8
                                         // s[0] = "b" doesn't work
                                         // let h = s[0] doesn't work
    for c in s.chars().rev() {
        println!("c = {}", c);
    }

    // if I really want the first character of the string
    if let Some(first_char) = s.chars().nth(0) {
        println!("first {}", first_char);
    }

    // String - heap allocated object
    let mut letters = String::new();
    let mut a = 'a' as u8;
    while a <= ('z' as u8) {
        letters.push(a as char);
        letters.push(',');
        a += 1;
    }
    println!("{:?}", letters);

    // &str <> String
    let u: &str = &letters; // magic from dref conversions

    // concatenation
    // String + str
    let z = letters + "abc"; // String + str
                             // z = letters + &letters; // String + String - do dref conversion
    
    let mut abc = String::from("hello world"); // or "hello_world".to_string();
    abc.remove(0);
    abc.push('!');
    println!("{}", abc.replace("ello", "goodbye"));
}

pub fn demo() {
    structures();
    enums();
    option(2.0f64);
    option(0.0f64);
    arrays();
    vectors();
    slices();
    strings();
}