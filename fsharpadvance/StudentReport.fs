namespace StudentReport

module Student =    
    
    type Student = {
        Name: string
        Marks: int list
    }

    let average student =
        student.Marks
        |> List.averageBy float

    let create name marks = 
        {
            Name= name
            Marks= marks
        }