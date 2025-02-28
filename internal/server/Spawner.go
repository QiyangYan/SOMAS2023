package server

import (
	"SOMAS2023/internal/common/objects"
	"SOMAS2023/internal/common/utils"

	baseserver "github.com/MattSScott/basePlatformSOMAS/BaseServer"
	"github.com/google/uuid"
)

const BikerAgentCount = 6
const AudiAgentCount = 0

func GetAgentGenerators() []baseserver.AgentGeneratorCountPair[objects.IBaseBiker] {
	return []baseserver.AgentGeneratorCountPair[objects.IBaseBiker]{
		baseserver.MakeAgentGeneratorCountPair[objects.IBaseBiker](BikerAgentGenerator, BikerAgentCount),
		baseserver.MakeAgentGeneratorCountPair[objects.IBaseBiker](AudiAgentGenerator, AudiAgentCount),
	}
}

func BikerAgentGenerator() objects.IBaseBiker {
	return objects.GetIBaseBiker(utils.GenerateRandomColour(), uuid.New())
}

func AudiAgentGenerator() objects.IBaseBiker {
	panic("not implemented")
}

func (s *Server) spawnLootBox() {
	lootBox := objects.GetLootBox()
	s.lootBoxes[lootBox.GetID()] = lootBox
}

func (s *Server) replenishLootBoxes() {
	for i := 0; i < LootBoxCount-len(s.lootBoxes); i++ {
		s.spawnLootBox()
	}
}

func (s *Server) spawnMegaBike() {
	megaBike := objects.GetMegaBike()
	s.megaBikes[megaBike.GetID()] = megaBike
}

func (s *Server) replenishMegaBikes() {
	for i := 0; i < MegaBikeCount-len(s.megaBikes); i++ {
		s.spawnMegaBike()
	}
}
