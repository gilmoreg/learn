open System

type CarType =
    | Tricar = 0
    | StandardFourWheeler = 1
    | HeavyLoadCarrier = 2
    | ReallyLargeTruck = 3
    | CrazyHugeMonsterTruck = 4
    | WierdContraption = 5

type Car(color: string, wheelCount: int)=
    do
        if wheelCount < 3 then
            failwith "Assuming 'cars' have at least 3 wheels"
        if wheelCount > 99 then
            failwith "Why so many?"
    
    let carType = 
        match wheelCount with
            | 3 -> CarType.Tricar
            | 4 -> CarType.StandardFourWheeler
            | 6 -> CarType.HeavyLoadCarrier
            | x when x % 2 = 1 -> CarType.WierdContraption
            | _ -> CarType.CrazyHugeMonsterTruck
    
    let mutable passengerCount = 0

    new() = Car("red", 4)

    member x.Move() = printfn "The %s car (%A) is moving" color carType
    
    // Like a property in C# (get only)
    member x.CarType = carType

    // As a "virtual" and with a setter
    // Default implementation required unless class is decorated [<AbstractClass>]
    abstract PassengerCount: int with get, set
    default x.PassengerCount with get() = passengerCount and set v = passengerCount <- v

type Red18Wheeler()=
    inherit Car("red", 18)

    override x.PassengerCount
        with set v =
            if v > 2 then failwith "only two passengers allowed"
                else base.PassengerCount <- v

[<EntryPoint>]
let main argv =
    let car = Car()
    car.Move()
    let greenCar = Car("green", 6)
    greenCar.Move()
    printfn "Green car has %i passengers" greenCar.PassengerCount
    greenCar.PassengerCount <- 1
    printfn "Green car now has %i passengers" greenCar.PassengerCount
    let truck = Red18Wheeler()
    truck.PassengerCount <- 2
    truck.Move()
    
    // Upcasting
    let truckObject = truck :> obj
    let truckCar = truck :> Car
    
    // Downcasting - when you don't know if the derived type is available
    let truckObjectBackToCar = truckObject :?> Red18Wheeler

    let result = System.Console.ReadKey()
    0 // return an integer exit code
