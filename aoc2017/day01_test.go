package aoc2017_test

import (
	"testing"

	"github.com/janreggie/AdventOfCode/aoc2017"
	"github.com/janreggie/AdventOfCode/structs"
	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	testCases := []structs.TestCase{
		{Input: "1122",
			Result1: "3",
			Result2: "0"},
		{Input: "1212",
			Result1: "0",
			Result2: "6"},
		{Input: "1234",
			Result1: "0",
			Result2: "0"},
		{Input: "91212129",
			Result1: "9",
			Result2: "6"},
		{Details: "Y2017D01 my input",
			Input: "6169763796227664136644229724736711773811471986347364813198244972" +
				"8688116728695866572989524473392982963976411147683588415878214189" +
				"9961635335845471757941581181487242988327988983333997865614591526" +
				"4414466995988734148196831917298735798978579136673284993278834377" +
				"2112176614723858474959919713855398876956427631354172668133549845" +
				"5856322119355736621813316131378698666932593743221698116836353253" +
				"2159724288935814712335811777491465378737136857478437672165218179" +
				"2371635288376729784967526824915192526744935187989571347746222113" +
				"6255779634761419231875346584456155969876143859115139392922572637" +
				"2351877488817463596325462476968453353144374572934434197374646932" +
				"6838186248448483587477563285867499956446218775232374383433921835" +
				"9931364633836288611155731428543589432911487662996536331955821359" +
				"3454496465766319838779444244353196461516965524365269678244339463" +
				"9169687847463721585527947839992182415393199964893658322757634675" +
				"2744229932379553541851948686384548914428939356944543242359681559" +
				"1396328264264996815328462615411147838991431676578343436545835278" +
				"5868895582488312334931317935669453447478936938533669921165437373" +
				"7414483784773918127799715289754782986887549392164214292517275555" +
				"9648194332226628952799667285638764867416699773134255898657525879" +
				"3261986817177487197512282162964167151259485744835854547513341322" +
				"6477326624435122518867718876516141776792299842711912923747559154" +
				"5737277585617853996513131956827825232624261515141277225425784741" +
				"3799811417287481321745372879513766235745347872632946776538173667" +
				"3712289772121439963916179743679234399237743885238455897693413511" +
				"6731139878779758354343472537434361172437939956619743215414688134" +
				"4528319826434554239373666962546271299717743591225567564655511353" +
				"2551975165152139638623837622589599574747897185647588433673257945" +
				"8988685241331471369891185518377897872255874232942986723926146477" +
				"3646389484318446574375323674136638452173815176732385468675215264" +
				"7367862428662956489973654126374996928177479379826285189263819392" +
				"7993599371241893856748828924677945843217933513973195216752752137" +
				"7546376518126276",
			Result1: "1182",
			Result2: "1152",
		},
	}
	for _, tt := range testCases {
		tt.Test(aoc2017.Day01, assert)
	}
}
