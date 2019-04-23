module Tests

open Xunit
open Learnfsharp.Calculator

[<Fact>]
let ``add returns correct sum`` () =
    Assert.True(Adder.add 1 2 = 3)

[<Fact>]
let ``add' returns correct sum`` () =
    Assert.True(Adder.add' 1 2 = 3)

[<Fact>]
let ``mult returns correct product`` () =
    Assert.True(Multiplier.mult 1 2 = 2)

[<Fact>]
let ``square returns correct square`` () =
    Assert.True(Multiplier.square 2 = 4)