package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val    int
	Key    int
	Height int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

func getMaxHeight(root *TreeNode) int {
	var leftHeight int
	var rightHeight int

	if root.Left != nil {
		leftHeight = root.Left.Height
	} else {
		leftHeight = 0
	}

	if root.Right != nil {
		rightHeight = root.Right.Height
	} else {
		rightHeight = 0
	}

	return int(math.Max(float64(rightHeight), float64(leftHeight)))
}

func setHeights(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		root.Height = 1
	}

	setHeights(root.Left)
	setHeights(root.Right)

	root.Height = getMaxHeight(root) + 1
}

func balanceTreeClockWise(root *TreeNode) *TreeNode {
	leftRoot := root.Left

	if leftRoot.Left == nil && leftRoot.Right != nil {
		rRightRoot := leftRoot.Right
		leftRoot.Right = rRightRoot.Left

		if rRightRoot.Left != nil {
			rRightRoot.Left.Parent = leftRoot
		}

		rRightRoot.Left = leftRoot
		root.Left = rRightRoot
		rRightRoot.Parent = root
		leftRoot.Parent = rRightRoot
		leftRoot = rRightRoot
	}

	root.Left = leftRoot.Right
	if leftRoot.Right != nil {
		leftRoot.Right.Parent = root
	}
	leftRoot.Right = root

	leftRoot.Parent = root.Parent
	if root.Parent != nil {
		if root.Parent.Left != nil && root.Parent.Left.Key == root.Key {
			root.Parent.Left = leftRoot
		} else {
			root.Parent.Right = leftRoot
		}
	}

	root.Parent = leftRoot

	return root
}

func balanceTreeCounterClockWise(root *TreeNode) *TreeNode {
	rightRoot := root.Right

	if rightRoot.Right == nil && rightRoot.Left != nil {
		lLeft := rightRoot.Left

		rightRoot.Left = lLeft.Right

		if lLeft.Right != nil {
			lLeft.Right.Parent = rightRoot
		}

		lLeft.Right = rightRoot
		rightRoot.Parent = lLeft
		lLeft.Parent = root
		root.Right = lLeft
		rightRoot = lLeft
	}

	root.Right = rightRoot.Left
	if rightRoot.Left != nil {
		rightRoot.Left.Parent = root
	}

	rightRoot.Left = root

	rightRoot.Parent = root.Parent
	if root.Parent != nil {
		if root.Parent.Left != nil && root.Parent.Left.Key == root.Key {
			root.Parent.Left = rightRoot
		} else {
			root.Parent.Right = rightRoot
		}
	}
	root.Parent = rightRoot
	return root
}

func balance(root *TreeNode) *TreeNode {
	var rightHeight int
	var leftHeight int

	if root.Left != nil {
		leftHeight = root.Left.Height
	} else {
		leftHeight = 0
	}

	if root.Right != nil {
		rightHeight = root.Right.Height
	} else {
		rightHeight = 0
	}

	if math.Abs(float64(rightHeight-leftHeight)) >= 2 {
		if rightHeight > leftHeight {
			root = balanceTreeCounterClockWise(root)
		} else {
			root = balanceTreeClockWise(root)
		}
		root = root.Parent
		setHeights(root)
	} else {
		root.Height = int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
	}

	if root.Parent == nil {
		return root
	} else {
		return balance(root.Parent)
	}
}

func insert(root *TreeNode, node *TreeNode) bool {
	if root.Key == node.Key {
		root.Val = node.Val
		return false
	}

	if root.Key > node.Key {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
			return true
		} else {
			return insert(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
			return true
		} else {
			return insert(root.Right, node)
		}
	}
}

func inOrderSuccessor(node *TreeNode) *TreeNode {
	successor := node.Right

	for successor.Left != nil {
		successor = successor.Left
	}

	return successor
}

func inOrderPredecessor(node *TreeNode) *TreeNode {
	predecessor := node.Left

	for predecessor.Right != nil {
		predecessor = predecessor.Right
	}

	return predecessor
}

func remove(root *TreeNode) *TreeNode {
	parent := root.Parent

	if root.Left == nil && root.Right == nil {
		if parent == nil {
			return nil
		} else {
			if root.Parent.Left != nil && root.Parent.Left.Key == root.Key {
				root.Parent.Left = nil
			} else {
				root.Parent.Right = nil
			}
		}
	} else if root.Right != nil {

		successor := inOrderSuccessor(root)
		root.Key = successor.Key
		root.Val = successor.Val

		if successor.Parent.Left != nil && successor.Parent.Left.Key == successor.Key {
			successor.Parent.Left = successor.Right
		} else {
			successor.Parent.Right = successor.Right
		}

		if successor.Right != nil {
			successor.Right.Parent = successor.Parent
		}

		parent = successor.Parent

	} else if root.Left != nil {
		predecessor := inOrderPredecessor(root)

		root.Key = predecessor.Key
		root.Val = predecessor.Val

		if predecessor.Parent.Left != nil && predecessor.Parent.Left.Key == predecessor.Key {
			predecessor.Parent.Left = predecessor.Right
		} else {
			predecessor.Parent.Right = predecessor.Right
		}

		if predecessor.Right != nil {
			predecessor.Right.Parent = predecessor.Parent
		}

		parent = predecessor.Parent

	}

	return balance(parent)
}

func search(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}

	if root.Key == key {
		return root.Val
	}

	if root.Key > key {
		return search(root.Left, key)
	} else {
		return search(root.Right, key)
	}
}

func searchNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Key == key {
		return root
	}

	if root.Key > key {
		return searchNode(root.Left, key)
	} else {
		return searchNode(root.Right, key)
	}
}

type MyHashMap struct {
	root *TreeNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (hm *MyHashMap) Put(key int, value int) {
	if hm.root == nil {
		hm.root = &TreeNode{value, key, 1, nil, nil, nil}
	} else {
		node := &TreeNode{value, key, 1, nil, nil, nil}
		needsBalance := insert(hm.root, node)
		if needsBalance {
			hm.root = balance(node)
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if hm map contains no mapping for the key */
func (hm *MyHashMap) Get(key int) int {
	return search(hm.root, key)
}

/** Removes the mapping of the specified value key if hm map contains a mapping for the key */
func (hm *MyHashMap) Remove(key int) {
	node := searchNode(hm.root, key)

	if node == nil {
		return
	}

	hm.root = remove(node)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	hm := Constructor()
	//operations := make([]string, 0)
	//operationsKeyVal := make([][]int, 0)

	//operations := []string{"put","remove","put","remove","put","remove","get","put","put","put","remove","remove","put","get","remove","put","get","remove","put","put","put","remove","put","get","put","remove","remove","remove","put","put","put","put","put","put","remove","put","get","get","remove","put","put","remove","put","put","remove","put","put","put","put","put","put","get","put","put","put","put","put","remove","put","put","put","put","put","put","put","put","put","put","put","remove","get","get","put","remove","remove","put","put","remove","put","remove","put","put","put","put","put","put","put","put","remove","put","put","put","put","put","put","remove","put","put","put","remove","put","get","put","remove","put","put","remove","put","put","put","get","put","put","put","put","get","put","put","put","get","put","put","put","put","put","get","remove","remove","put","put","put","put","put","put","put","put","remove","put","remove","put","put","put","remove","put","put","put","put","put","put","put","put","get","remove","put","put","remove","put","put","put","remove","put","put","put","remove","put","put","remove","put","put","put","put","remove","get","put","put","put","put","put","put","put","put","remove","put","put","put","put","get","put","remove","put","put","remove","put","put","put","put","put","put","put","remove","put","put","put","put","put","get","put","put","get","put","put","remove","put","put","put","put","put","put","put","put","put","put","put","remove","remove","put","put","put","put","put","put","put","put","remove","get","put","put","get","put","put","put","put","get","remove","put","put","put","put","put","remove","remove","remove","put","put","put","put","get","get","remove","put","get","put","put","put","put","put","put","put","get","put","put","put","put","put","get","put","remove","put","put","put","put","remove","put","put","put","get","put","put","put","remove","remove","put","put","put","put","get","put","put","put","put","get","get","put","put","put","put","put","put","put","put","get","put","put","put","put","get","put","put","put","get","put","get","put","get","put","get","put","put","get","put","put","put","put","put","put","put","put","put","put","put","put","put","get","put","get","get","put","put","remove","remove","remove","put","put","put","put","put","get","put","put","put","put","remove","get","put","put","remove","put","put","get","put","put","put","put","put","put","put","put","put","put","remove","put","put","put","put","remove","put","remove","put","put","get","put","put","remove","put","get","get","get","put","get","put","put","put","put","put","put","put","put","remove","put","get","put","put","get","put","put","put","put","put","put","remove","put","remove","put","remove","remove","get","put","put","put","put","put","put","put","put","get","put","put","put","put","put","put","get","put","put","get","put","remove","get","put","get","put","remove","put","put","get","remove","put","get","put","get","remove","put","put","put","put","put","put","get","put","put","put","put","put","put","put","put","put","remove","put","put","put","put","put","put","put","put","get","put","put","put","get","put","put","put","remove","put","put","put","get","remove","put","put","put","put","get","get","put","get","put","get","put","put","put","put","get","get","put","put","put","put","put","put","put","put","put","put","put","remove","put","remove","put","put","get","remove","put","get","put","put","put","put","put","put","remove","put","get","put","put","put","put","put","put","put","get","put","put","put","put","put","put","put","put","remove","put","put","put","put","put","put","put","put","put","put","put","put","put","remove","get","put","put","put","put","put","put","put","put","put","remove","get","put","put","put","put","put","put","get","put","get","remove","put","put","get","put","get","remove","put","get","put","put","put","put","put","get","get","put","remove","put","put","remove","put","put","get","put","get","get","remove","put","put","put","put","put","put","put","remove","put","remove","get","put","put","get","put","put","put","remove","put","put","get","put","put","put","put","put","remove","put","remove","get","remove","put","put","remove","get","get","get","put","remove","put","remove","put","remove","put","put","put","get","put","put","put","put","put","put","put","remove","put","put","put","put","put","put","put","remove","put","put","remove","put","put","get","put","put","put","get","get","put","put","put","put","put","put","get","get","put","put","remove","put","put","put","get","put","get","get","put","remove","get","put","put","put","get","put","get","put","remove","get","put","put","remove","get","put","put","put","get","put","put","put","put","put","put","get","put","put","put","put","put","put","put","get","put","put","put","get","put","remove","remove","get","put","put","remove","get","remove","get","put","put","put","put","put","put","put","remove","get","get","remove","put","remove","put","remove","put","remove","put","put","get","remove","put","put","put","put","get","put","put","put","put","remove","put","put","get","remove","put","put","put","remove","remove","remove","remove","put","remove","put","put","put","remove","put","put","put","put","remove","put","get","get","get","put","put","put","put","put","put","remove","remove","put","put","put","remove","put","put","put","put","remove","put","put","put","put","put","put","put","put","put","put","remove","get","put","get","put","remove","put","put","put","get","put","put","put","put","put","remove","remove","put","put","put","put","remove","put","put","remove","get","put","get","remove","put","put","put","get","put","put","get","get","put","put","put","put","get","put","put","put","put","get","put","get","put","put","remove","put","put","remove","put","put","put","put","put","remove","put","remove","remove","put","put","remove","remove","put","put","remove","put","remove","put","put","put","put","put","put","put","put","remove","put","put","put","put","put","remove","get","get","put","remove","put","put","remove","get","get","put","get","put","put","remove","remove","put","put","get","remove","remove","put","get","put","remove","put","put","get","get","put","put","get","put","put","put","put","put","get","put","remove","put","put","put","put","put","put","get","put","get","put","put","put","get","put","get","put","put","get","put","put","get","put","put"}
	//operationsKeyVal := [][]int{
	//{769,729},{769},{379,724},{415},{421,11},{507},{421},{217,686},{815,1},{328,330},{379},{217},{288,43},{769},{655},{656,1},{769},{656},{321,821},{625,812},{675,867},{379},{202,993},{516},{550,289},{421},{772},{769},{11,88},{377,773},{67,462},{671,376},{896,507},{555,200},{157},{86,420},{11},{118},{613},{562,346},{649,603},{328},{386,277},{25,587},{86},{528,682},{784,536},{444,8},{120,730},{703,701},{262,339},{520},{739,158},{256,235},{667,217},{648,804},{428,264},{256},{388,115},{456,737},{400,253},{537,343},{169,899},{520,354},{813,591},{627,558},{651,122},{127,650},{236,515},{550},{909},{456},{547,331},{283},{781},{255,422},{185,116},{547},{806,454},{981},{34,271},{190,749},{371,319},{409,688},{572,871},{561,913},{926,136},{933,655},{933},{679,600},{129,415},{747,348},{642,57},{779,9},{246,254},{377},{29,796},{59,142},{9,520},{649},{371,102},{59},{699,184},{679},{986,418},{80,805},{377},{848,525},{94,559},{921,773},{236},{336,493},{345,949},{272,357},{184,851},{806},{658,36},{148,897},{15,250},{328},{740,758},{911,41},{718,854},{34,673},{942,168},{869},{769},{671},{612,746},{342,509},{786,421},{22,595},{165,409},{50,865},{118,300},{504,245},{115},{425,726},{550},{801,390},{31,947},{507,782},{184},{631,373},{277,959},{207,54},{572,739},{459,809},{966,824},{567,576},{596,146},{377},{127},{753,184},{96,82},{146},{123,813},{731,647},{939,247},{625},{467,65},{333,151},{655,893},{336},{397,452},{180,226},{813},{745,450},{436,394},{409,771},{414,802},{240},{732},{175,942},{533,623},{6,988},{490,539},{74,315},{85,599},{166,330},{593,570},{848},{987,575},{637,332},{746,874},{583,731},{31},{679,701},{740},{688,404},{431,263},{108},{424,825},{724,331},{498,178},{411,867},{316,872},{532,531},{113,681},{166},{85,888},{371,23},{425,595},{716,703},{992,147},{6},{943,308},{733,471},{333},{9,570},{874,239},{642},{198,318},{825,560},{675,283},{374,852},{44,781},{674,258},{785,301},{160,179},{206,242},{48,982},{514,334},{942},{878},{643,916},{774,554},{732,183},{320,611},{698,412},{343,157},{163,772},{980,835},{94},{779},{392,129},{737,401},{324},{188,145},{286,898},{293,533},{546,505},{379},{561},{650,73},{773,45},{740,745},{482,617},{805,528},{897},{89},{179},{841,925},{78,675},{816,328},{701,434},{493},{679},{747},{478,820},{141},{578,201},{241,998},{685,287},{30,953},{538,307},{594,818},{591,852},{345},{42,447},{379,17},{957,886},{496,593},{243,435},{504},{209,493},{371},{152,697},{449,276},{896,954},{884,463},{822},{338,706},{221,589},{790,970},{968},{440,665},{649,696},{662,829},{733},{160},{891,935},{641,409},{263,935},{34,79},{862},{70,368},{707,221},{702,829},{579,928},{707},{908},{430,878},{997,253},{125,250},{650,665},{563,104},{747,757},{396,401},{720,547},{557},{935,426},{623,74},{288,286},{324,905},{351},{920,48},{778,724},{927,432},{731},{225,582},{263},{533,903},{190},{515,486},{272},{129,370},{635,259},{733},{129,995},{160,730},{166,234},{489,206},{274,14},{487,210},{381,974},{252,976},{111,710},{460,45},{777,320},{516,66},{270,315},{982},{637,7},{463},{328},{343,417},{221,466},{710},{507},{869},{708,928},{408,96},{371,17},{206,267},{898,283},{770},{895,821},{590,469},{901,860},{260,855},{85},{513},{81,935},{847,245},{180},{261,518},{410,197},{108},{316,998},{110,564},{853,846},{902,894},{159,763},{782,782},{343,132},{361,470},{548,692},{374,810},{797},{49,952},{88,341},{475,396},{394,849},{812},{137,375},{338},{491,842},{791,583},{489},{291,44},{226,690},{671},{822,627},{895},{190},{942},{711,81},{641},{280,538},{44,493},{387,782},{788,375},{568,127},{548,334},{881,746},{763,348},{59},{549,747},{926},{130,639},{716,805},{676},{732,46},{803,26},{672,713},{366,62},{601,290},{970,635},{348},{569,81},{243},{66,115},{725},{217},{779},{245,77},{465,620},{76,573},{336,868},{607,977},{282,399},{975,132},{507,375},{821},{618,256},{719,718},{74,233},{777,493},{984,130},{799,237},{578},{573,932},{31,578},{206},{29,812},{821},{731},{47,656},{740},{778,350},{835},{710,661},{182,217},{126},{483},{41,279},{224},{903,187},{785},{223},{514,848},{269,891},{66,640},{661,464},{146,832},{331,94},{867},{148,398},{710,224},{375,602},{880,245},{707,274},{964,97},{751,385},{791,627},{231,342},{386},{138,198},{75,702},{368,191},{750,257},{206,26},{339,212},{3,644},{820,7},{74},{915,665},{224,396},{659,232},{152},{70,70},{146,278},{954,369},{337},{67,852},{276,36},{67,916},{149},{853},{427,846},{126,731},{72,939},{52,813},{745},{449},{566,879},{522},{478,428},{911},{871,883},{805,733},{541,516},{256,779},{661},{420},{942,169},{341,892},{850,982},{869,920},{929,319},{264,468},{797,342},{710,620},{468,831},{949,887},{459,85},{887},{490,132},{926},{764,899},{64,888},{524},{537},{784,834},{96},{416,927},{965,668},{593,861},{566,923},{606,205},{502,883},{733},{764,621},{716},{873,23},{834,463},{74,426},{755,568},{106,132},{34,567},{841,86},{863},{616,693},{310,949},{37,923},{156,963},{951,894},{922,10},{163,290},{0,741},{597},{467,607},{533,3},{900,526},{745,726},{146,728},{545,312},{228,677},{986,707},{469,573},{385,975},{343,270},{518,515},{633,450},{577},{712},{351,24},{999,369},{86,823},{462,628},{143,954},{164,600},{282,872},{967,337},{930,439},{428},{801},{206,997},{84,453},{45,360},{568,627},{132,610},{743,452},{790},{812,52},{733},{583},{465,226},{172,437},{999},{536,984},{636},{929},{931,618},{942},{977,473},{201,192},{792,501},{166,900},{688,177},{823},{205},{106,405},{25},{44,676},{962,593},{861},{18,235},{500,244},{876},{101,167},{979},{244},{409},{449,567},{195,373},{102,188},{232,737},{360,327},{282,164},{183,1},{820},{124,918},{635},{666},{740,208},{780,189},{949},{544,212},{240,92},{222,670},{316},{560,352},{580,993},{448},{846,767},{35,399},{567,991},{753,397},{567,715},{858},{36,645},{269},{146},{41},{683,888},{599,830},{797},{753},{531},{80},{565,401},{519},{924,190},{797},{605,182},{387},{121,324},{290,792},{586,479},{611},{887,298},{222,162},{871,67},{4,225},{142,693},{228,895},{873,300},{651},{857,205},{712,432},{562,472},{810,954},{429,326},{78,43},{564,622},{328},{671,594},{814,186},{803},{138,348},{416,65},{633},{118,986},{775,479},{231,444},{579},{277},{9,454},{355,256},{642,568},{902,303},{653,722},{653,507},{627},{872},{290,251},{332,133},{351},{704,661},{189,807},{336,331},{44},{806,702},{548},{903},{761,914},{796},{774},{985,531},{353,238},{419,278},{193},{852,764},{565},{408,901},{357},{911},{816,630},{303,711},{70},{110},{739,774},{123,594},{392,950},{411},{326,952},{447,175},{2,707},{56,664},{678,553},{280,673},{733},{106,44},{751,148},{961,791},{974,625},{166,277},{854,167},{711,1},{590},{742,286},{451,373},{472,327},{671},{312,61},{834},{980},{753},{841,417},{165,194},{444},{482},{394},{874},{94,474},{709,631},{337,40},{203,126},{575,185},{159,702},{805,324},{992},{315},{914},{962},{774,958},{180},{988,737},{590},{643,514},{254},{788,272},{40,288},{255},{78},{503,656},{475,215},{801,114},{972,211},{672},{303,949},{282,824},{49,781},{177,420},{578},{422,712},{71,470},{129},{490},{73,720},{69,810},{301,847},{282},{955},{646},{148},{946,109},{217},{113,84},{193,63},{821,321},{254},{53,884},{317,741},{436,353},{301,858},{410},{653,415},{642},{29},{126},{557,808},{416,903},{174,479},{908,308},{70,858},{925,8},{443},{69},{180,550},{989,999},{812,912},{106},{392,538},{874,661},{952,421},{136,554},{312},{314,361},{323,444},{800,154},{262,995},{906,910},{521,844},{903,45},{459,90},{154,737},{984,772},{869},{562},{174,941},{518},{602,525},{970},{716,98},{795,901},{414,605},{887},{697,421},{650,543},{719,450},{575,467},{633,997},{822},{754},{160,90},{483,920},{527,946},{643,895},{627},{806,264},{578,755},{360},{937},{547,9},{814},{572},{837,604},{911,162},{690,181},{649},{571,561},{561,93},{154},{565},{375,150},{548,348},{87,215},{305,873},{394},{63,109},{366,6},{832,633},{716,129},{646},{742,274},{935},{884,285},{912,720},{48},{29,404},{150,181},{825},{336,912},{64,689},{667,438},{6,358},{656,37},{953},{100,549},{180},{286},{531,2},{993,474},{11},{314},{850,585},{162,464},{910},{773,298},{484},{494,150},{135,182},{755,816},{843,96},{34,737},{228,754},{485,112},{638,128},{341},{118,422},{550,536},{624,802},{414,599},{575,602},{342},{775},{975},{923,88},{419},{900,806},{174,69},{155},{217},{228},{512,123},{75},{365,561},{525,808},{310},{516},{429,575},{516,493},{150},{832},{502},{339,818},{256},{749,247},{969},{697,848},{991,680},{816},{641},{322,465},{720,181},{639},{34,787},{404,939},{629,20},{703,7},{304,496},{924},{210,155},{814},{305,82},{740,404},{502,69},{254,907},{953,215},{404,240},{339},{951,482},{791},{309,136},{504,762},{955,940},{185},{834,856},{483},{594,915},{325,838},{780},{748,957},{696,438},{992},{998,587},{78,414}}

	operations := []string{"remove", "put", "remove", "remove", "get", "remove", "put", "get", "remove", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "remove", "put", "put", "get", "put", "get", "put", "put", "get", "put", "remove", "remove", "put", "put", "get", "remove", "put", "put", "put", "get", "put", "put", "remove", "put", "remove", "remove", "remove", "put", "remove", "get", "put", "put", "put", "put", "remove", "put", "get", "put", "put", "get", "put", "remove", "get", "get", "remove", "put", "put", "put", "put", "put", "put", "get", "get", "remove", "put", "put", "put", "put", "get", "remove", "put", "put", "put", "put", "put", "put", "put", "put", "put", "put", "remove", "remove", "get", "remove", "put", "put", "remove", "get", "put", "put"}
	operationsKeyVal := [][]int{{27}, {65, 65}, {19}, {0}, {18}, {3}, {42, 0}, {19}, {42}, {17, 90}, {31, 76}, {48, 71}, {5, 50}, {7, 68}, {73, 74}, {85, 18}, {74, 95}, {84, 82}, {59, 29}, {71, 71}, {42}, {51, 40}, {33, 76}, {17}, {89, 95}, {95}, {30, 31}, {37, 99}, {51}, {95, 35}, {65}, {81}, {61, 46}, {50, 33}, {59}, {5}, {75, 89}, {80, 17}, {35, 94}, {80}, {19, 68}, {13, 17}, {70}, {28, 35}, {99}, {37}, {13}, {90, 83}, {41}, {50}, {29, 98}, {54, 72}, {6, 8}, {51, 88}, {13}, {8, 22}, {85}, {31, 22}, {60, 9}, {96}, {6, 35}, {54}, {15}, {28}, {51}, {80, 69}, {58, 92}, {13, 12}, {91, 56}, {83, 52}, {8, 48}, {62}, {54}, {25}, {36, 4}, {67, 68}, {83, 36}, {47, 58}, {82}, {36}, {30, 85}, {33, 87}, {42, 18}, {68, 83}, {50, 53}, {32, 78}, {48, 90}, {97, 95}, {13, 8}, {15, 7}, {5}, {42}, {20}, {65}, {57, 9}, {2, 41}, {6}, {33}, {16, 44}, {95, 30}}
	//random := rand.New(rand.NewSource(time.Now().UnixNano()))
	//
	//for i := 0; i < 1000; i++ {
	//  operations = append(operations, "put")
	//  operationsKeyVal = append(operationsKeyVal, []int{random.Intn(100000), random.Intn(100000)})
	//}

	for index, operation := range operations {
		switch operation {
		case "remove":
			key := operationsKeyVal[index][0]
			fmt.Printf("removing key %d\n", key)
			hm.Remove(key)
		case "get":
			key := operationsKeyVal[index][0]
			fmt.Printf("value for key %d, is %d\n", key, hm.Get(key))
		case "put":
			key := operationsKeyVal[index][0]
			value := operationsKeyVal[index][1]

			fmt.Printf("Putting %d, %d\n", key, value)
			hm.Put(key, value)
		}
	}
}
