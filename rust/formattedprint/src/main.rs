// Import (via `use`) the `fmt` module to make it available.
use std::fmt::{self, Formatter, Display, Result};

#[derive(Debug)]
struct Person<'a> {
    name: &'a str,
    age: u8
}

// Define a structure for which `fmt::Display` will be implemented. This is
// a tuple struct named `Structure` that contains an `i32`.
#[derive(Debug)]
struct Structure(i32);

// To use the `{}` marker, the trait `fmt::Display` must be implemented
// manually for the type.
impl Display for Structure {
    // This trait requires `fmt` with this exact signature.
    fn fmt(&self, f: &mut Formatter) -> Result {
        // Write strictly the first element into the supplied output
        // stream: `f`. Returns `fmt::Result` which indicates whether the
        // operation succeeded or failed. Note that `write!` uses syntax which
        // is very similar to `println!`.
        write!(f, "{}", self.0)
    }
}

struct List(Vec<i32>);

impl Display for List {
    fn fmt(&self, f: &mut Formatter) -> Result {
        // Extract the value using tuple indexing,
        // and create a reference to `vec`.
        let vec = &self.0;

        write!(f, "[")?;

        // Iterate over `v` in `vec` while enumerating the iteration
        // count in `count`.
        for (count, v) in vec.iter().enumerate() {
            // For every element except the first, add a comma.
            // Use the ? operator to return on errors.
            if count != 0 { write!(f, ", ")?; }
            write!(f, "{}", v)?;
        }

        // Close the opened bracket and return a fmt::Result value.
        write!(f, "]")
    }
}

fn main() {
    let name = "peter";
    let age = 27;
    let peter = Person { name, age };

    println!("{:#?}", peter);

    println!("{subject} {verb} {object}",
             object="the lazy dog",
             subject="the quick brown fox",
             verb="jumps over");

    println!("{} of {:b} people know binary, the other half doesn't", 1, 2);

    println!("{number:>width$}", number=1, width=6);
    println!("{number:0>width$}", number=1, width=6);


    println!("Compare debug to display");
    let structure = Structure(32);
    println!("Display: {}", structure); // output: Display: 32
    println!("Debug: {:?}", structure); // output: Debug: Structure(32)

    let v = List(vec![1, 2, 3]);
    println!("{}", v);

    // Exericses
    let pi = 3.141592;
    println!("pi is approximately {:.3}", pi);

}
