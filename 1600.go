
type ThroneInheritance struct {
	dead map[string]bool
	kingName string
	child map[string][]string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		dead: make(map[string]bool),
		kingName: kingName,
		child: make(map[string][]string),
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.child[parentName] = append(this.child[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	ans := make([]string, 0)
	var dfs func(root string)
	dfs = func(root string) {
		if d,ok := this.dead[root]; ok && d {
			// ignore
		} else {
			ans = append(ans, root)
		}
		for _, child := range this.child[root] {
			dfs(child)
		}
	}
	dfs(this.parentName)
	return ans
}