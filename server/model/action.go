package model

type ActionType string

const (
	ActionMoveForward   ActionType = "ACTION_MOVE_FORWARD"
	ActionTurnLeft      ActionType = "ACTION_TURN_LEFT"
	ActionTurnRight     ActionType = "ACTION_TURN_RIGHT"
	ActionMoveUp        ActionType = "ACTION_MOVE_UP"
	ActionMoveDown      ActionType = "ACTION_MOVE_DOWN"
	ActionDig           ActionType = "ACTION_DIG"
	ActionDigUp         ActionType = "ACTION_DIG_UP"
	ActionDigDown       ActionType = "ACTION_DIG_DOWN"
	ActionSuck          ActionType = "ACTION_SUCK"
	ActionSuckUp        ActionType = "ACTION_SUCK_UP"
	ActionSuckDown      ActionType = "ACTION_SUCK_DOWN"
	ActionGetGPS        ActionType = "ACTION_GET_GPS"
	ActionDrop          ActionType = "ACTION_DROP"
	ActionDropUp        ActionType = "ACTION_DROP_UP"
	ActionDropDown      ActionType = "ACTION_DROP_DOWN"
	ActionCraft         ActionType = "ACTION_CRAFT"
	ActionDropIfType    ActionType = "ACTION_DROP_IF_TYPE"
	ActionRecipe        ActionType = "ACTION_RECIPE"         //Hmmm - could be too complex as an Action
	ActionSelect        ActionType = "ACTION_SELECT"         //Hmmm - could be too complex as an Action
	ActionMoveInventory ActionType = "ACTION_MOVE_INVENTORY" //Hmmm - could be too complex as an Action
)

type Action struct {
	ActionType     ActionType
	DigIfBlocked   bool
	RetryUntilDone bool
}
