Non-standard Golang kit for [Lux Season 2](https://github.com/Lux-AI-Challenge/Lux-Design-S2).

## Notes on main.py and JSON format

* Kaggle expects to use `main.py` as the entry point, which also has to have some magical properties relating to an `agent()` function
* Here we just use a copy of the JavaScript kit's `main.py` - modified to call the bot instead of `node main.js`
* Note that the Go code expects the incoming JSON messages to have the complete (non-sparse) arrays provided
* The JavaScript kit's `main.py` does this (most other kits' `main.py` do not)

## Notes on submissions

* Kaggle is running Linux
* Submissions need to be `tar.gz` files containing `main.py` and the compiled file bot
* The bot file at least needs its file permissions set (e.g. to 0o777)
* A Python script is provided which allows Windows users (like me) to automate the whole process

## Notes on internal data structures

* The main data structure is the `Frame` type, which is simply the JSON message sent by `main.py`, unmarshalled
* Since the structure of that is fairly complex, `Frame` is complex too - see `structs.go`
* **NOTE THAT OBJECTS INSIDE THE FRAME ARE NOT VALID IN LATER TURNS** - get the current object from the current frame instead
* Some helper methods are included to quickly get needed items, e.g. `GetBoard()`, `MyFactories()`, etc

## Pointers vs values

* The `Pos` type should be stored/passed by value
* Then you can compare one `Pos` to another by simple `==` checking
* Also, if you get a `Pos` from a `*Unit`, it will be safe to mutate it if you want to
* `*Unit` and `*Factory` objects should be stored/passed by pointer, because:
* The API stores new actions in the objects themselves
* These are then sent to the `lux_s2` referee at the end of each turn
* If you make copies of the objects, any actions in them won't be seen
* So just keep pointers to them instead, then you've got the right ones

## API outline (possibly subject to change, WIP)

```golang
func Run(bidder func(*Frame), placer func(*Frame), main_ai func(*Frame))

func (self *Frame) Bid(faction string, bid int)
func (self *Frame) PlaceFactory(pos Pos, metal int, water int)
func (self *Frame) RealStep() int 
func (self *Frame) GetBoard() *Board
func (self *Frame) GetCfg() *EnvCfg
func (self *Frame) Width() int
func (self *Frame) Height() int
func (self *Frame) MyUnits() []*Unit
func (self *Frame) TheirUnits() []*Unit
func (self *Frame) AllUnits() []*Unit
func (self *Frame) MyFactories() []*Factory
func (self *Frame) TheirFactories() []*Factory
func (self *Frame) AllFactories() []*Factory
func (self *Frame) FactoryByStrain(n int) *Factory
func (self *Frame) FactoryAt(xy XYer) *Factory
func (self *Frame) CanPlaceFactory() bool
func (self *Frame) PotentialSpawns() []Pos
func (self *Frame) RandomSpawn() Pos

func (self *Unit) BuildRequest(args ...Action)
func (self *Unit) ClearRequest()
func (self *Unit) HasRequest() bool
func (self *Unit) PowerAfterRequest() int
func (self *Unit) CanAcceptRequest() bool
func (self *Unit) IsMine() bool
func (self *Unit) NaiveTrip(other XYer) []Action

func (self *Factory) Act(action FactoryActionType)
func (self *Factory) ClearRequest()
func (self *Factory) HasRequest() bool
func (self *Factory) IsMine() bool

type XYer interface {		// Implemented by Pos, *Unit, and *Factory
	XY() (int, int)
}

func Dist(a XYer, b XYer) int
```
