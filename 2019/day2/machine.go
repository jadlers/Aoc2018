package main

type machine struct {
	memory []int
	ip     int
}

func (m *machine) add() bool {
	if !m.safeToAccessMemory(m.ip+1, m.ip+3) {
		return false
	}

	i1 := m.memory[m.ip+1]
	i2 := m.memory[m.ip+2]
	t := m.memory[m.ip+3]
	m.memory[t] = m.memory[i1] + m.memory[i2]
	m.ip += 4

	return true
}

func (m *machine) mul() bool {
	if !m.safeToAccessMemory(m.ip+1, m.ip+3) {
		return false
	}

	i1 := m.memory[m.ip+1]
	i2 := m.memory[m.ip+2]
	t := m.memory[m.ip+3]
	m.memory[t] = m.memory[i1] * m.memory[i2]
	m.ip += 4

	return true
}

// Run executes the instructions in memory until the exit code
func (m *machine) Run() (ok bool) {
	ok = true
	for ok {
		op := m.memory[m.ip]

		switch op {
		case 1:
			ok = m.add()
		case 2:
			ok = m.mul()
		case 99:
			return
		}
	}

	return false
}

func (m *machine) safeToAccessMemory(start, end int) bool {
	for _, position := range m.memory[start : end+1] {
		if position < 0 || position >= len(m.memory) {
			return false
		}
	}
	return true
}
