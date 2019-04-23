namespace Learnfsharp
open Learnfsharp.Calculator

module Program =
    [<EntryPoint>]
    let main argv =
        printfn "Hello World from F#!"
        printfn "square of 5 is %i" (Multiplier.square 5)
        printfn "sum of 1 and 2 is %i" (Adder.add 1 2)
        printfn "sum of 2 and 3 is %i" (Adder.add' 2 3)
        let exit = System.Console.ReadLine()
        0 // return an integer exit code
