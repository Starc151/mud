package player

import (
	"fmt"
)

func Start() {
	const startLocation int = 13
	hp, mp := PlayerStart()
	fmt.Println("Привет! Вы в игре!")
	Action(startLocation, hp, mp)
}
