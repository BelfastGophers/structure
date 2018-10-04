# structure

Two examples of structuring a Go application based off a talk at the Belfast Gophers October 2018 meetup. 

[Slides used in talk](https://docs.google.com/presentation/d/1QWS6yItYxWjDxxCOdWBueffHonT7BCZzdPUeF0q8AuU/edit?usp=sharing)

## mvc

This is close to a standard mvc layout but with domains split to separate files and a flat dependency structure.

Good to mid size projects but after 15 or so domain objects it can become a bit harder manage and browse as the files are split across separate packages.

## domain-driven

Is closer to a DDD style layout with each domain having it's own package with handlers, service & data access siloed in each.
