package smartlogging

type Level struct {
	name  string
	value int
}

type LevelsRepository struct {
	TRACE   Level
	DEBUG   Level
	INFO    Level
	WARNING Level
	ERROR   Level
	FATAL   Level
}

// Logger is a struct that allows to store all the information
// relative to the current logger.
type Logger struct {
	level Level
}

// Levels is a container of all the levels
var Levels LevelsRepository

func init() {
	Levels = LevelsRepository{
		Level{"TRACE", 0},
		Level{"DEBUG", 10},
		Level{"INFO", 20},
		Level{"WARNING", 30},
		Level{"ERROR", 40},
		Level{"FATAL", 50},
	}
}
