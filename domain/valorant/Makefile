generate-abilities:
	protoc --go_out=../../ ability.proto

generate-agents:
	protoc --go_out=../../ agent.proto

generate-maps:
	protoc --go_out=../../ map.proto	

generate-weapons:
	protoc --go_out=../../ weapon.proto

generate-all:
	make generate-abilities
	make generate-agents
	make generate-maps
	make generate-weapons