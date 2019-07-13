#![allow(dead_code)]
#![allow(unused_variables)]
use std;
use std::mem;

struct Point {
  x: f64,
  y: f64,
}

fn origin() -> Point {
  Point { x: 0.0, y: 0.0 }
}

pub fn stack_and_heap() {
  // to use heap instead of stack
  let y = Box::new(10);
  println!("Box {}", *y);
  // let z = vec![10, 20, 30];
  // println!("vec! {}", z);


  // stack allocated
  let p1 = origin();
  // heap allocated
  let p2 = Box::new(origin());

  println!("p1 {} bytes", mem::size_of_val(&p1));
  println!("p2 {} bytes", mem::size_of_val(&p2));

  // now on stack
  let p3 = *p2;
}