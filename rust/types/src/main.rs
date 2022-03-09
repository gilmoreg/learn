
use std::f64::consts;
use std::mem;

mod control_flow;

mod data_structures;
mod stack_and_heap;

// const requires type, has no fixed address, is just inlined during compile
const MEANING_OF_LIFE: u8 = 42;

// static
static Z: i32 = 123;
// static mut will require unsafe blocks to access
static mut MUTABLE_Z: i32 = 123;

fn main() {
    fundamental_data_types();
    operators();
    scope_and_shadowing();
    stack_and_heap::stack_and_heap();

    println!("MEANING_OF_LIFE {}", MEANING_OF_LIFE);
    println!("Z {}", Z);
    unsafe {
        println!("MUTABLE_Z {}", MUTABLE_Z);
    }
    control_flow::demo();
    data_structures::demo();
}

fn scope_and_shadowing() {
    // any curly braces open a scope
    let a = 14;
    {
        // shadowing
        let a = 12;
        println!("{}", a); // 12, not 14
    }
    println!("{}", a); // 14, not 12

}

// Tuples can be used as function arguments and as return values
fn reverse(pair: (i32, bool)) -> (bool, i32) {
    // `let` can be used to bind the members of a tuple to variables
    let (integer, boolean) = pair;

    (boolean, integer)
}

fn operators() {
    let mut a = 12 * 4 + 2;
    a = a + 1; // ++/-- not supported
    a += 1; // this is supported though
    a %= 7; // mod
    let a_cubed = i32::pow(a, 3);
    println!("{} {}", a, a_cubed);

    let b = 2.5;
    let b_cubed = f64::powi(b, 3);
    let b_to_pi = f64::powf(b, consts::PI);
    println!("{} {} {}", b, b_cubed, b_to_pi);

    // bitwise
    let c = 1 | 2; // 01 | 10 = 11 = 3_10
    let two_to_ten = 1 << 10;
    println!("{} {}", c, two_to_ten);
}

fn fundamental_data_types() {
    // unsigned - cannot be negative
    let a: u8 = 123; // 8 bits, 0-255
    println!("a = {}", a);
    // a = 456; cannot change immutable variables

    // mutable, signed
    let mut b: i8 = 0;
    println!("b = {}", b);
    b = -42;
    println!("b = {}", b);

    let mut c = 123456789; // 32-bit signed i32
    println!("c = {}, size = {} bytes", c, mem::size_of_val(&c));
    c = -1;
    println!("c = {}, size = {} bytes", c, mem::size_of_val(&c));

    // isize/usize
    let z: isize = 123;
    let size_of_z = mem::size_of_val(&z);
    println!("size of z {} bytes", size_of_z * 8);

    let d: char = 'x';
    println!("d = {}, size = {} bytes", d, mem::size_of_val(&d));

    let e = 2.5; // double precision, 8 bytes, f64, can use f32 for floats if necessary
    println!("e = {}, size = {} bytes", e, mem::size_of_val(&e));

    let g = false;
    println!("g = {}, size = {} bytes", g, mem::size_of_val(&g));
}