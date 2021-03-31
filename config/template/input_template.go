package template

type InputTemplate struct {
	Id 				string								 `yaml:"id,omitempty"`
	Title 			string                               `yaml:"title"`
	Global 			bool                                 `yaml:"global"`
	Name 			string                               `yaml:"name"`
	Configuration   *InputConfigurationTemplate 		 `yaml:"configuration"`
	Node 			string                               `yaml:"node,omitempty"`
	AutoStart 		bool                            	 `yaml:"auto_start"`
}

func (it *InputTemplate) Equals(other InputTemplate) bool {
	return it.Title == other.Title &&
		it.Global == other.Global &&
		it.Name == other.Name &&
		it.Node == other.Node &&
		inputConfigurationPointerEquals(it.Configuration,other.Configuration)

}

/*
From: http://web.mnstate.edu/peil/MDEV102/U1/S6/Complement3.htm

Complement of a Set of Inputs: The complement of a set of Inputs, denoted diff,
is the set of all elements in the given universal set a that are not in b.
diff = {x ∈ a : x ∉ b}.
*/
func InputComplement(a, b []InputTemplate) (diff []InputTemplate) {

	if len(b) == 0 {
		return a
	}

	for i := range a {
		for j := range b {
			if a[i].Equals(b[j]) {
				break
			} else if j + 1 == len(b) {
				diff = append(diff,a[i])
			}
		}
	}
	return diff
}
