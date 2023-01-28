Very raw, non-standard Golang kit for [Lux Season 2](https://github.com/Lux-AI-Challenge/Lux-Design-S2)

## Notes on main.py and JSON format

* Kaggle expects to use `main.py` as the entry point, which also has to have some magical properties relating to an *agent()* function
* Here we just use a copy of the JavaScript kit's `main.py` - modified to call `./bot` instead of `node main.js`
* Note that the Go code expects the incoming JSON messages to have the complete (non-sparse) arrays provided
* The JavaScript kit's `main.py` does this

## Notes on submissions

* Kaggle is running Linux
* Submissions need to be *tar.gz* files containing `main.py` and the compiled file `bot`
* The `bot` file at least needs its file permissions set (e.g. to 0o777)
* A Python script is provided which allows Windows users (like me) to automate the whole process

## Notes on internal data structures

* The main data structure is the *Frame* type, which is simply the JSON message sent by `main.py`, unmarshalled.
* Since the structure of that is fairly disorganised, *Frame* is disorganised too.
* Some helper methods are included to quickly get needed items, e.g. *GetBoard()*, *MyFactories()*, etc.
