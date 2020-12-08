package passwordphilosophy

import (
	"testing"
)

func TestPasswordsValidNumExample(t *testing.T) {
	list := []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
	result := passwordsValidNum(list)
	expected := 2

	if result != expected {
		t.Errorf("When getting the number of valid passwords, got: %v, want: %v.", result, expected)
	}
}

func TestPasswordValidNumLong(t *testing.T) {
	list := []string{"3-4 l: vdcv", "6-9 d: dddddkzdl", "6-13 f: mfswqfrqffrvfvf", "10-12 l: sfzlnwcptlnlflq", "2-4 m: qbwcmt", "15-16 v: vvvvvvvvnvvvvcvvvvgv", "1-4 n: wnnzfln", "1-3 x: xxgx", "7-8 j: jjjjjfvh", "6-8 x: xxzxxnnwx", "3-8 t: djtdznbtgwtrhxf", "7-9 w: glwwwwxtxwlwwwcp", "4-5 g: jggjtgggg", "9-10 c: trgjrcbfcwf", "9-10 v: vvvbvmvcvvvv", "3-14 p: pptzpppppppppsfrsppc", "3-5 n: nnznmkn", "1-13 t: fxtrpgtzztsfnmzmmtg", "13-15 d: dcztmqjdmclgdhdcd", "3-4 j: mjjjjc", "13-18 x: fxxgxxxxxxxxxxxxxcx", "2-4 f: fxqs", "3-4 w: ffmbwwnn", "12-13 k: dlpdzkhbkkkftkb", "5-11 r: vrdxrrrrrtrqz", "7-8 b: bbbbbbrcb", "13-16 c: wzccvzccclccczcc", "10-11 m: lmmmmmmmmwqmm", "6-10 f: fffqjfznvfhz", "8-9 h: hvfhhvgklhhpfmh", "4-17 p: hpprnpsppcppppppfp", "9-13 z: zzzzrgtzzzzzzzzzzl", "6-10 l: lllllqlllplll", "3-10 j: vjjjwjjjjjjjj", "4-8 j: jjjjjjjjjjc", "2-5 s: jsgtsfsw", "1-2 b: brwpbbbbb", "2-5 m: tbkpfbvnqmpjbbx", "7-8 m: xxbgmbmnmmwmdhmm", "9-12 v: phsslrznvrsvtxgv", "10-14 v: rnnvvrvvvzbhmcx", "13-15 b: bbbbbbbbbbgbwbbb", "2-5 w: bwldjwwz", "5-7 d: dddjthk", "10-12 c: pckcccccccccl", "2-11 b: rdjqkdbzxzqfgn", "3-8 s: dsssmshs", "8-10 d: dddfdddsdj", "2-4 l: slgltjzwzrnlllrh", "6-10 z: vrcvgrgzphxd", "10-13 b: mnblbbbbbbbfvbbbbx", "8-16 n: wznnnbfhndnnnnngnn", "6-8 v: fvvvzqks", "3-6 n: znqnbnzwnn", "15-16 g: wggzxsgcgjwgxhgggg", "10-13 x: xlvxxxxxkxxxfxx", "2-6 z: jpznbgszmj", "13-19 j: jjzjjjdjwxlnjhwsrtj", "5-6 p: bvppqqppcp", "19-20 r: rrrrrrrrbrrrrrrrrrvd", "4-7 j: jjjjjjljj", "5-6 k: kkkkkck", "1-2 k: zbkk", "11-13 r: rrsrrrwrrrrrr", "6-11 z: zzgzznzzzzwzzrzpzgc", "7-11 l: llvlllllllllllll", "7-8 f: fffffjffnp", "3-4 j: pxcw", "9-11 k: zkkkkkkkdkzx", "12-13 g: fgglgvggxhgks", "6-7 d: kbdddftdddl", "14-15 m: dmmjwnxxfdfhwvcdd", "3-4 b: bbbsbt", "2-9 d: tgsvznmnnrgvtldjbj", "4-5 l: ljxtc", "5-8 j: jnvsjlfw", "3-4 v: zrklbf", "1-15 v: vvvvvvvvvvvvvvtv", "5-6 n: ssnnnngnsxkphtnwrk", "6-7 k: kkkkkkdk", "2-4 d: xchmdk", "12-13 v: vfvvmvvvpvvwvpvvvdm", "3-14 d: dzcddwdddddfdfdwd", "1-6 b: dbbbbsb", "12-14 b: bbbxbbbbbbbjbb", "1-8 s: sssssssgn", "5-10 d: ddddwddddvdd", "10-12 s: sssmssssskswssss", "3-13 q: cdqfcqvszbzqnhvxlw", "8-10 z: tmgzkvxzzvz", "8-10 n: nnnnntnqnhrsnn", "2-7 r: vjrctrhprhc", "6-8 w: bzfxrwwfmx", "3-5 x: xxxwxb", "5-9 z: vbspkzfsrrpdcfl", "8-9 z: jgqzzzzzwq", "5-12 x: dvqvjsdfjxgtxkd", "5-6 r: rrrrrr", "1-4 f: xdzngd", "15-18 x: xxxxxxxxxxxxxxvvxx", "1-4 g: qcwjpxgtvgkf", "13-15 m: mhmmmmmmmhmmqmmmmmmv", "5-11 n: nnnndnnnnnnnnnnnn", "4-9 n: tjgnpqsnpxpzfw", "10-13 q: qqlqdqqqqqqqnqq", "8-14 g: gdcggnknjsgdskq", "3-5 l: vlllll", "16-18 t: ttttttttttttttttmt", "9-14 b: kbhbcbbbbbqbsbr", "3-9 f: jzrnxffljfmhffftdng", "1-4 z: pbzs", "10-11 h: lrrgpmhshtfg", "1-5 n: mnqjwjpfnnzn", "1-8 r: zrrrprrkr", "8-9 g: gkbtggcpf", "3-8 q: qqsqqgqqqqm", "17-18 j: jjjjjjjjjjjjjjjjgkj", "13-14 t: tttttttttmtttfttt", "2-6 c: cvcccqc", "4-17 c: sfnckrcmlcpmcpcgbkct", "10-12 d: dddddddqvddd", "3-5 d: ddddtd", "10-11 h: hchhhhhhhvg", "3-11 j: mmlcjsjhjbpjjjxsk", "11-12 w: wwwwwwwwwwwww", "9-10 c: ccccccccccc", "5-7 x: ckhdnfgjrdkxqphvjtvs", "10-12 s: ssssjsssshstd", "7-8 m: mmmmmmmg", "3-4 v: vjvvt", "2-3 f: hmknf", "3-5 g: bgsglggg", "12-13 x: vxxxdcgxxxfxx", "3-20 h: hhhtkvrmfswvmkzmbvmh", "2-3 v: vvvv", "3-14 g: zgpggzgqxjzggg", "15-16 c: ggcjccnccccccccc", "11-12 d: wrdrdxsfdzddkl", "6-7 q: vqqknqq", "2-6 w: wwwwwwwkck", "4-6 l: dplldljlvxrrm", "2-18 m: pmwfkqlknsxzjmdzkmxd", "5-11 q: rqkzqxqmqtqnzqqqxk", "6-7 s: sssssskss", "9-10 w: wwwwwwwwsw", "5-7 x: jxxxhxwxxxxxxxx", "4-5 t: ttbttfm", "2-4 h: mlhhhkhc", "13-14 l: llllllllllllfjl", "8-10 k: kllkkkjkkkk", "1-3 s: sssbdss", "2-3 z: zlbzzzzzz", "9-10 r: rrvrrrfrdfrr", "4-12 k: vkkkbkhkzkkk", "11-12 l: llllllllllcgc", "5-6 k: hkzkkkkgdsghdk", "3-5 n: nzsxk", "12-14 r: rwrrrrrrrrrprxrrr", "4-7 k: kkxkkkkkkk", "12-16 b: bbbbbbbbbbbbbbbbb", "13-16 p: ppppppppppppzppfp", "15-19 j: tjjjjjdjjjjmjjccbjm", "1-3 l: llll", "3-4 x: hxkl", "9-15 n: ktnpmnnnlnxnnft", "6-7 q: qqdtcqs", "13-14 g: lgghndddgllnfg", "7-14 k: cdzkjkkksksbwkjb", "7-11 w: cwcsvqwfwhw", "14-18 m: mmmmmmmmmmmmmhmmmjmm", "4-5 c: cvfctcsncscbmj", "4-5 r: zrnrrnxct", "2-3 b: ggdlb", "8-9 x: cpxskxxxxbxn", "12-14 v: vvvvvwvvvvvbvw", "13-16 n: pmhnhpnntncxqrmd", "2-4 m: mmxm", "7-10 m: mmmrqwsmnkgbdbrmmjjm", "7-8 z: zzzzzzzzzz", "2-3 p: pppnfltzpd", "1-10 h: hhslhfsgtxphhncg", "5-6 m: txmhxhtmgm", "4-5 v: wsnvrsvx", "8-12 p: pfbdlnrphdfplvpqx", "10-13 b: nbpjvgghsbplb", "9-15 h: ghcxbptxsgvvvptqx", "4-5 t: tttgf", "2-3 l: lsnl", "1-11 v: vvvvvvvvvvvvvvvv", "6-7 r: brnrrrr", "3-4 k: bkkvkrtmzj", "17-18 p: pppppwphpkpppppppp", "8-11 f: fqfffdffffpfffffffff", "14-16 z: zznzzzbzzzzdzzzsz", "10-11 v: xtvfnxdrvqnxg", "8-13 c: tcwtbcpfvdccs", "5-6 b: btbnbdb", "2-4 d: ddxmcb", "6-9 p: qpppzppbpsplpppppxp", "7-12 z: zzzzzztzqzzqzz", "4-5 h: hhhmb", "7-10 f: frftffhfff", "4-5 c: xccmv", "8-11 x: rfglxxxhxwvxxfcxx", "7-10 k: kkqnjxzckmmbkkk", "6-7 h: hhxsdhhwh", "15-18 r: rmrrlrrrrrrrrtncrcr", "10-11 x: srnxxbbhxxx", "13-14 c: cccccccccccchx", "5-8 r: rrrrfrrhrrr", "11-15 v: vvvvvvvvvvvvwvvvsvv", "6-8 q: qqkqlnqzmqqx", "2-3 h: hhzq", "10-12 l: lllflllllpqt", "7-9 m: mlmmmmmmm", "4-12 b: pbbhbrrhzxvrgcsfmqq", "8-10 g: tggrgggbgdggmg", "9-10 s: sssssksspsss", "4-8 k: dtkkkkkkkkkkmk", "19-20 j: jjjjjjjjjjjjjjjjjjjj", "3-8 r: rrrrxrrrrrrr", "9-10 p: qmgzfbqpnp", "1-4 w: wwwsw", "1-15 f: fffffffftzftxhfwfbx", "1-5 t: rtttttlxtjtttsq", "3-4 f: fmffkt", "16-18 n: kfnnnlnnnsgscnzmrv", "3-4 w: wwmk", "7-9 l: lllkbllllllllljtjll", "7-13 l: pqzbtxmqslllqkx", "6-18 j: qlnjjvrtzqkfrnjxjk", "6-8 q: qtpqqpqg", "6-8 c: cccfccxd", "3-4 f: fhff", "4-5 x: xjxfhqkhx", "16-18 k: xvbktrmqkckwkrkkjkk", "9-13 v: vvvbvvvvvvvvnvvlmvv", "1-7 k: wkkkkkqkk", "2-4 z: jpzlxfjbh", "5-8 l: lllllpllfllllr", "7-8 r: rrrrrrwnr", "5-6 d: ddddftd", "4-5 q: vqlqn", "3-13 p: pppsppmpbppppppppn", "1-3 w: wwww", "4-5 d: tfrpfcdlft", "5-10 r: rrrrppxrrvrm", "9-10 t: dtttttttgvtt", "8-12 j: jjjhjjjljjjjjpjj", "6-7 b: lpfpzljhbl", "3-6 k: kzkqkrkckkv", "3-7 q: rvqdtsqqjfftzllflnx", "2-3 l: glsjcsrlmn", "5-8 g: gggxhggpgg", "12-14 q: zqqsqqqqqqqqxqqq", "9-11 q: qqqqsqjtqqq", "9-10 g: gggggcggggg", "12-14 z: zzzzzzzzzzzlzzzzzz", "1-6 s: gvqttcgrqqsbndv", "8-19 z: mzzzpkzdzzzskfhzzlj", "4-6 h: kmlbhwthhp", "8-9 x: xxxxxxxvk", "1-2 k: drjkkkk", "10-12 z: zdzsgzzzwvzrzszzzzzz", "13-14 f: lffdffjfbfdxjd", "11-12 b: rppkbnjgcqrr", "1-2 l: lllq", "5-6 m: mmmmzr", "6-14 l: lllbldllpllllblllcl", "4-9 d: rdtdcwnbqjbthddr", "13-15 v: vvcwvvqvzvvvfvg", "9-10 g: gpxpwxhmgg", "13-14 n: njpndvsvfgrrhnb", "3-5 r: rrrrr", "7-13 h: hhdcgscdxmhwhg", "3-4 l: tnhf", "10-13 c: mfcvbcmcccqhc", "6-7 x: cdwfdxxxxxx", "5-7 b: wbbzwbcx", "4-6 h: vhhzpqjhk", "5-8 r: hffprngr", "14-15 p: pppppppppppspjm", "8-12 z: pzqdhzzkzzzzxzzc", "1-8 b: kbbfdvvcsdjlbw", "16-18 k: rqwtgtwzgsnvlkdkkk", "8-10 w: xllfslwwsf", "2-5 m: qmbdm", "10-12 b: bhbbbbbwtbbbbbhbb", "8-14 c: cscdccccccccxcp", "3-5 c: rcdpz", "16-18 q: qqqqqqqqqqnqqqqbql", "3-4 x: xjxx", "6-7 n: gnnnnxwqn", "16-17 f: tffrffffffffffffff", "5-6 d: ddddgm", "2-7 r: prrrssrrkrjrrr", "12-16 g: ggxbxggtgbggccwgggtv", "10-14 n: nnkntpnvncnnnznnr", "3-5 c: qccwcx", "7-17 v: jvvvvhvvvpvvqvgvv", "3-6 c: cclccs", "8-10 m: nwtmfcnmmmmm", "1-2 p: hrpp", "2-3 j: hdnjgt", "1-3 h: jhgvhjhh", "9-10 j: jjjjjjjjvl", "12-15 h: hhshhhhthhhhhnhn", "5-8 k: qkvmkwckvlkl", "12-15 j: jjjjjjjkjjjwjjlj", "11-16 v: tvmrmvvvwwqrxwvg", "5-7 b: dbhtpjvxqzbhbz", "3-4 b: mqbg", "4-7 t: tttbtvt", "8-13 h: hhhhchhhhqfhhhhhthr", "1-6 w: lhwqgrbbwbv", "1-11 t: ttttpttttttt", "2-5 l: rlwllvfhw", "7-8 m: jmmrmmmmlmmm", "3-5 s: smsss", "15-19 p: mfpzcghppcgpgqpfvxfw", "6-7 n: jnldsnb", "7-9 x: xwbwdxjhcxxvxx", "1-2 c: cbqprgfjfpc", "3-4 g: gngf", "7-10 w: twvvwwtwwtwwsgwws", "2-9 g: ggngggglgrzk", "7-9 p: xpphpkppppp", "11-12 p: pppppxpppppppppp", "3-4 g: fzqrl", "6-7 x: xxxxxmw", "3-4 r: rrrn", "8-11 b: bbhbbbbrbbwbkb", "3-11 d: xxvfmvrdmzhfmvmhhp", "1-12 z: mzzzkzwzzzxfzzzzzzzz", "5-11 g: gzbjgxghddggkgq", "5-8 g: grpggrgg", "1-6 f: lfffffsfffff", "8-10 x: slxxxgxxqxxx", "4-13 l: lllllllllllllvll", "8-9 z: zzzzdhzzz", "13-15 l: rglrmcqvgjzvlklx", "2-4 q: nsxs", "2-3 v: vvqcv", "13-14 b: bbbbbbbbbbbbbb", "13-14 x: rxxdxnxxtxxxqq", "1-4 s: jmsssssgssvppss", "6-7 j: pgvjkxcvjslmj", "2-5 w: jwwswgl", "5-7 m: mmnmmxp", "2-6 k: tkkkxwtkrk", "4-15 n: nnnnngnnnnnnnnnnnt", "7-11 x: dxxtxxxvwtnx", "11-14 s: ssrvphskdsstzsszrnsz", "11-12 l: rclllgbllblb", "6-7 x: fxhbmnj", "15-17 q: kqvhwsqhqqqhmqqtqqlq", "12-14 k: kksktkkkkkkkkfk", "3-4 h: hshk", "14-17 v: vvvqvvkvvvvvvvjkvxv", "2-6 f: vfvfrp", "4-12 z: bzzqszhwffrhr", "2-9 b: sbdjfbndbhcclrmblccn", "11-13 j: jjjjjjjjjjgjkjjjjj", "15-17 f: fpffffsfffffffdfjfff", "6-7 v: tlvvvvv", "11-19 l: llllllllllllllllllll", "1-6 w: wwbwwwwww", "3-7 s: cpsvmss", "2-11 r: rrpnkchfzljwxmrrvvw", "7-12 n: msgpplgpfnkmxlm", "4-13 f: sjpfqpnzqzznfftfhgfs", "13-14 q: qqqqqqqqqqqqqqq", "13-14 j: jjjmjjjjjjjjgq", "2-4 q: qdrds", "5-10 s: ssssqssssbs", "6-8 c: ccdcchchcc", "18-20 z: xztgqzzckzgcpzzzjzdz", "5-8 z: ctpdzzzzzqzmvzxn", "6-8 b: bbmkgbfkb", "2-10 x: xxxxxxxxxxx", "5-6 z: gzzzqk", "4-7 c: cccvccj", "7-8 d: dddddddd", "9-11 c: ccdccfccxcccc", "1-4 z: zjzzzz", "6-8 z: zlmjbzvz", "8-10 v: gvtjvgvzvjphscn", "8-10 d: dcdddddddddddddd", "7-8 p: mjpwppppp", "1-11 w: wwwwwnwwwwwww", "7-9 j: jjjjjjjjj", "6-8 q: dmqdqqzt", "3-10 p: tsqcpgmzgh", "11-15 s: ssslsxrlbfqsssssss", "5-10 m: dpmtmhbdmk", "3-5 p: wpjqv", "10-13 p: lhvkzxprppmtprj", "4-6 w: nlfjwz", "19-20 c: kzckcfsqpdvssgcgktcm", "1-2 p: ppgltbv", "3-4 w: ngwwfttwqg", "10-11 p: ppppppnpllgp", "2-4 t: tkttptt", "3-6 x: xwrjpcf", "14-17 q: qcqbqfqqkqdqrqqqq", "4-8 n: qsncdnbknpn", "8-10 s: wzscpmssrshh", "3-8 p: pfppppgpp", "11-15 b: tqdbgfbbbgjbbbg", "12-15 h: lwhhjzhpmhzhhnhmhh", "4-5 w: nwccwlpwdvtlwwr", "11-16 w: wwwwwwwwwwswwwwxw", "6-9 q: qqqqqqqqqqqqqqqq", "8-12 t: tqttjtttpgbtlt", "14-15 v: vvvvkvgvvvvvvvg", "2-12 f: fpfflxzxffwxfvfff", "7-12 m: hmmmmwmmbmlm", "2-3 z: mmlzzz", "12-16 h: hhbhhhkhhfhpbghm", "15-18 s: ssssxsssssmsssqssx", "8-9 s: sqsqdtmssqscxxs", "1-10 n: nnbnnngnnn", "10-14 c: ccbcvscncrccng", "7-8 k: kkhkxrkkctkpjkkck", "10-11 g: bgpgggkcwzs", "9-17 l: llllllllckwllllllwl", "7-12 q: gqqqjbxtqvqxzjhvjrjj", "1-3 s: ssnsc", "3-10 n: nnnnnnnnnnnnnnnn", "9-12 l: blfnqllsllll", "5-8 s: sgsgstssl", "1-4 c: cccwcc", "7-9 h: hhhhhhhhhhh", "2-4 s: sssss", "8-10 f: vznpfffrsffscgfzft", "4-5 r: hnrns", "4-5 z: zzzlgz", "10-18 w: wwwwwwwwwwwwwwwwwwww", "3-12 b: bfmbbbbfbbltbg", "14-15 z: zzjxrhrjsbrrzgd", "1-2 z: zzzzz", "4-11 x: pxgxwphxmgxpt", "1-16 c: cbgnxcccxsrcbmcvqc", "3-6 n: cnppnnjnxnnjss", "8-9 r: zltfbjwrrlmtrzqrh", "1-3 z: zzqzzzzzzz", "4-7 n: vgndgdndbxcnznxc", "2-9 w: wwwwwcfnww", "1-4 b: tbbrbbbbbcbbbbbbbbdk", "7-8 h: hhhwhzhhhhtf", "12-13 j: jjjjjjjjjjjjj", "2-8 m: rthcgnxmm", "2-7 l: glzctbjh", "5-7 d: dddddddx", "12-13 x: xxxxxxxxxxxxs", "6-11 t: ltnfnnqttvtkjnlk", "6-7 x: qxxxxpfxxx", "1-3 d: kdndzd", "6-11 f: zfxfmffptsffmffvfz", "1-4 x: xxxxxxx", "4-5 l: lllllllll", "4-8 c: cjccccgckrdpvv", "13-14 z: zzzwzzzzzzzzkkzp", "1-3 r: lgtr", "3-17 g: ggtgjgmggggggggggggg", "4-5 q: qqqqqq", "2-6 v: vvvlvvv", "4-17 c: nwqcprgftnqgcrjpcjb", "7-10 x: xlxfrjxxxxgnxx", "2-14 q: qqqqqqqqqqqqqsqqq", "4-8 x: xxxxxxxxxxx", "1-16 z: bzzzzzzzzmzzzzldzzzz", "2-4 v: bffhzd", "3-4 r: rrrr", "5-8 x: xxxxxnsxd", "1-17 g: gbngjggxgggjgzvjk", "2-3 b: sgbb", "11-14 p: ppppppppppppppp", "3-9 q: vhsjtqppqgdqdqlf", "1-3 w: wcwzsfzwwcwdwncwp", "7-11 f: ffffffdfffff", "3-8 g: ksgdcqqgtkmgjbzsf", "5-8 s: tsssccswss", "1-7 r: brrrrrhrrrrr", "5-7 g: ggxmgggcwgz", "2-4 h: hhwh", "9-18 w: wwwwwwwwhwwwwwwwwbw", "1-7 m: ndmmbmmmmmmm", "5-7 j: jjvjvjt", "3-8 m: plnmcmkfmmcm", "4-5 r: rrrjzr", "4-8 k: kkkgkjknkjkk", "17-18 s: srngdqpwxcrbqrprss", "4-9 t: mbkjrtntztrttt", "1-10 w: hwwwwwwwwwwtx", "5-9 f: lfvhfxndfffmxhf", "7-8 n: pnnnnnbvnnnnnnnnnhnn", "1-10 b: jjtptpmmqv", "1-3 p: rqppqt", "1-11 t: pbtwdttkttmtxj", "5-6 g: jpkfggtcmj", "6-8 f: fffffpfvf", "9-13 t: ttttttttdtttg", "3-5 r: prrjr", "18-19 b: bbbbbbbbbbbbbbbbbbb", "5-7 v: vjvjxjhvrv", "8-10 z: zzzsznzxttgspzzzz", "5-13 m: mngmcbmqbxmzjfm", "1-3 b: xbhb", "12-14 m: mfmmbgmmmmmcmz", "3-6 n: jpnpnnn", "3-14 j: xmgjqkzfctctxwgjms", "7-14 k: lbkfcbctkgkjkqvjdwkc", "6-8 l: lhthfgllfvbpxchct", "5-6 h: hhhhhk", "6-9 c: qkxtbqltkzcrdcgckncc", "13-15 w: fhsvgcxwfcgxwvwwwkb", "3-4 c: cccmcj", "3-4 m: mhmcmm", "4-5 k: skkqg", "4-16 m: mgmcmmmmmmmmmmmtmmm", "3-5 g: gggbsg", "9-15 l: dnslgllwlflmlllllll", "3-4 h: mhhh", "3-5 z: vkjzlzzvrztwg", "11-12 q: qqqqqqqqqqqq", "5-10 p: xpxprppppgzppp", "1-4 z: kzzhzszzzzzzzz", "2-7 j: nsjfszt", "8-13 j: qtrjfmbjfswrj", "2-10 s: ssscwcsnsssrt", "1-12 w: zwwwwwwwwwwx", "2-4 p: ppgvf", "2-3 n: hnnxmgtchqwgnfx", "11-12 z: bfzlvgdzpfzzzgvhz", "1-6 d: czqqdjrxs", "6-7 r: rrrrrkz", "4-7 d: hxdzddvq", "2-4 g: slwgs", "13-14 f: mqfhrlftvvfkff", "12-14 d: ddddddddddhqdc", "1-3 f: ffzff", "4-5 r: mrrrkrrrrrrkr", "3-4 m: mmwl", "2-5 j: jljjsjdj", "4-5 f: fkfdf", "16-17 s: dzdsrsgmsjffsnssls", "10-17 f: fffffhffffhffffff", "8-10 v: fvvvvvvvvz", "4-9 n: nnnpsnjstl", "8-11 x: xxzdxfjczxx", "4-10 k: nrkxhfzspmfpzcl", "7-10 x: pxxjlcbtsjxdc", "1-14 s: fsszssssssfsssss", "5-7 d: ddgdgdvddd", "2-11 t: jttpshfcmlt", "1-3 v: pvdv", "5-7 x: xxkxxxx", "13-14 d: ddddxdvdppdcrdhk", "2-6 h: hhmhpc", "6-11 g: jxhsggmgcrp", "13-16 r: krrrnccjrrtrtrrlprrr", "3-4 f: ffff", "5-7 t: ttttttttttttt", "5-7 d: ndddjdw", "7-12 w: wgjwwbwcxfww", "2-6 s: sxsssbs", "12-14 l: xlxslmmjpvmzll", "4-6 h: hdhhhhdhkch", "9-11 j: jjjjjjjjjjj", "6-8 d: dddddddn", "5-6 b: bbbbsnbbbbb", "1-8 h: hhhhhhhhhh", "5-11 p: ppppppppppp", "11-12 d: dddddddddddk", "1-2 b: bbpbb", "2-8 j: jjhhjzjjt", "5-6 k: kfkjkk", "3-5 j: kgjjjjqnjcfmjjrc", "2-5 l: jtmzhlcs", "3-6 v: vvvvvvv", "15-16 v: vpvvvvvvvvjvvsvvvvvv", "15-16 h: hhhhhhhhmhhhdhhh", "2-3 k: ksdg", "7-11 b: dbhbbbbbhbbrb", "4-14 k: kkkknkkkkkbtkkk", "5-6 j: smdndsgfm", "2-3 c: sccc", "10-15 w: bwwwbwwwcwwlwwpwwwww", "3-6 w: wwwwwww", "14-19 s: jqdgjszsssssdrssjszs", "5-10 p: vpppbppppp", "4-7 p: zcpltzbsjmcgpv", "2-8 j: tfjjjjjnprj", "14-15 m: mmmmmmmmmmmmmhqm", "1-6 t: pttttct", "9-10 n: qnbnzxjfnnkb", "5-6 v: vvvvjvv", "7-8 v: qvvbvpvvvf", "2-15 s: sjsssssssssssstss", "7-10 b: bvbbbbnbbbbb", "13-14 b: bqbbbjqbbbbbtq", "3-7 k: cltbmkkb", "8-9 c: mlscsmqccfmccszsslck", "3-4 k: kklckk", "5-7 c: cccmclg", "5-7 d: dnddddd", "1-5 f: xfffbffffffffff", "13-15 n: nnnnnnnnnnnnnnqfnn", "3-6 t: ttttqztttcttt", "3-4 s: ssnws", "9-11 f: bcvkzrdrcffgzjgvf", "3-10 r: rrrdrrqrrrrrr", "3-6 h: bbqhkh", "3-8 n: nngnwthw", "1-11 x: ffqbbnxcpkrlzhm", "3-4 b: bbbbnnb", "5-6 f: lfdwqjffmj", "7-16 v: vvvvvvvvvvvvvvvvv", "13-15 m: zmmmzmmmmmmmmkmrmmmm", "1-2 r: rrrr", "12-14 d: dkdddddddddddcddd", "14-15 r: crlrtjkrrnhqrrr", "5-9 w: xjjwpcpfffchxtww", "15-16 d: jdcchsgdjtjdcxdrlwx", "9-15 p: cpmpppprpnpppsp", "6-7 w: jlqkwww", "1-12 p: wpfqzppptpflp", "4-14 x: xxxxxxxxxxxxxxxxxx", "1-4 z: xzzjz", "2-10 t: zvbnlwltbkvvcf", "5-9 v: vpvvvvfvvvvv", "2-12 l: lllllllllllllllll", "2-6 w: wmcwzpvxqsbxrrw", "11-12 l: lldlllllllrllgll", "9-13 h: hhhhhhhhdhxhrhh", "6-9 x: xxxzxxxxxxxx", "1-15 x: vcsxxxgxxxxwxxbbxx", "7-12 d: fmddcdfndddddkdnrgt", "12-14 x: wlmbpdxxrxbxdxg", "5-9 l: gnllscnsz", "4-5 j: jjjjj", "6-10 q: kvctnxpqhqfhjb", "6-8 q: vqqnqcqpqq", "2-9 k: mvxsbrxzr", "10-11 c: lbcgbpsxhcc", "9-17 x: xkdpklsvxqlkvrlvx", "3-14 j: trjcmjzjqkbjjj", "7-8 k: kkkkkkhfg", "1-3 f: pfdf", "9-12 s: glcjzdmllprrrxmdds", "9-11 b: bwbbbbbbbbqb", "2-14 f: kzfffbzfhfffflf", "5-7 n: ngnnnnn", "6-10 t: tttttrtttjtt", "7-8 f: fvffzfpff", "3-5 h: jbhkx", "9-11 b: qckbnjtbcbb", "12-15 m: mmmmkmmmmqmmmmjxm", "9-11 c: cdckczpcccc", "1-3 v: wzxq", "8-11 q: qhqqqqqlqqnqq", "4-9 j: mpjjqsdsj", "13-17 s: sssssssssssssssss", "4-7 x: knbxqrxxzl", "15-17 m: mmmmmmmmmmmmmmmmmmm", "6-10 w: wxnwfhwqdd", "9-11 r: crrrrcrnrrwrrmrrrm", "9-15 l: klflxlcllldllllll", "1-4 n: vpnznn", "2-4 k: cjkg", "7-8 m: xqmmrmmjjshmfm", "4-6 f: sfffff", "2-7 l: mltmmnldsnnl", "5-6 j: tjjdjxjjbs", "4-5 p: ppmnp", "2-3 b: bzrb", "5-17 h: dmpdgxthrcppfznbhchw", "3-4 p: pppp", "3-5 v: vvhvvr", "9-10 g: gflzqwbvgsgnbzngmgr", "8-17 z: xpvzsrhgzqtdcfjzl", "1-4 x: khxjxxvd", "3-4 z: czzz", "10-12 m: zjzcmmpxmjxmxzm", "7-9 k: kkkkkkkktkzkkk", "16-17 d: ddrdgdxdddddgddwmd", "2-16 f: xmsflgsvfwnfmxwnnmg", "4-12 b: ltbkbqbbbbbc", "12-14 t: tctlnvtjttsgztthtxtt", "9-14 x: xxxgpxvlkxdxdrxxrxwx", "15-17 s: xsssssssssssssvgw", "10-16 q: hqqpqxjqbsxqfqbm", "4-5 v: vvvnvvvttv", "15-16 v: rwshvltmvgvcvpvv", "4-5 j: jjjjjjjj", "3-5 w: jwfwp", "4-5 m: nddmm", "10-13 t: rtttnttttttttj", "14-17 l: llllllllllllltlcnl", "7-16 p: crjhtmpppgkpvgnpt", "6-9 b: bbfbbhbfgdrb", "1-2 q: qkqq", "17-20 f: hwcftkmtzhftmnfwfsdf", "2-4 w: wjtww", "10-11 p: pppppppppppp", "5-12 l: dntvlnljwkkldgp", "3-4 s: gmkvlqxsx", "8-10 m: qmmmmmmmmm", "12-14 p: gtbnfmlvtkppwp", "7-8 t: ttttttvs", "12-18 j: jjjjjjjjjjjjjjjjjb", "5-7 p: tfrppwpcpppp", "11-13 h: hthhhhzfbhhhdhhph", "1-7 n: ncgsvwvlvwlhbtmnnnpx", "10-11 j: djjbsjljjxf", "5-12 n: cnjqkknnnnmjnq", "18-20 z: zzzzzzzzzzzzzzzzzbzp", "2-5 m: kphbm", "12-13 g: tzrqwrjgzkxgg", "8-9 w: mwwwkwwwww", "2-6 j: jbxjjzjjjjj", "9-10 j: jjjjjjjjznj", "5-6 j: jjjspjj", "8-18 g: jggmgpfggbvggvgggggl", "3-5 n: nnxnh", "3-18 n: nnlnnnnnngnnnnnnnhnn", "6-10 r: rcrctghspmrrbrcrrrrq", "18-19 s: sfxqssssgssscxsdrss", "7-10 b: pbbbbbbgvbzbbz", "15-17 m: mmmmmmmmmmmmmmhmmm", "4-11 s: sssgssssssssssssxfss", "7-9 g: ggtgggggg", "4-8 w: wwwwwwwt", "17-20 p: pppppppppppppppppppn", "2-5 w: hdxrw", "9-18 c: cccccccccwcccccccc", "1-9 l: qllllllllllllll", "15-16 g: ggggggggggggggnj", "8-9 s: sjszkrhsss", "12-18 g: ggggggggggggggggggg", "5-6 s: ksnsxspsw", "4-8 p: pppbpppcppppp", "4-10 f: fffffffffff", "1-2 d: dqdd", "7-8 w: wwjwvhvcwbwww", "3-5 h: xghfhzhqw", "4-5 q: qqqzpw", "5-9 q: fxkrhqrjq", "2-4 m: mmjm", "1-4 p: wpmg", "7-13 n: nnnqfdtnnnnftnlnnn", "2-11 w: vwphpjskljxgqsgqzph", "2-4 w: wwwwgfkmpcf", "6-10 w: htgtkwwrmwvgrdmkzt", "3-4 s: sssstss", "11-20 l: dqlqmkllqclbmzklllrt", "1-6 g: gfgcwgvzkrhgjslg", "2-13 j: jjjszjxjljjjjqjjvjjh", "4-19 q: qqqlqqqqqqqqqqqqqqfq", "11-14 j: jbjjjjjjjjrjjp", "10-11 f: fpffvfffspvkf", "4-5 v: vhvlvhgh", "6-8 h: nsfklqhtxg", "2-3 z: czwnzzp", "1-10 c: wcmvcrcmjmcmbcckfcdc", "2-4 f: hzbrjqqntbbffbfl", "2-5 b: bzmbwztrb", "7-9 z: bscfzxzxz", "13-14 z: zzzzznzzzzzzbn", "3-7 t: rlxttrcmfttlt", "4-6 h: rldcbhx", "11-12 g: vfgfdgffbrrg", "1-10 s: ssssssssjsss", "7-13 k: dqbklkhvlgphjqhvgc", "8-18 x: xxxxdxxxxxxxxxxxxzxl", "1-3 x: xxxx", "11-12 z: ttzhvzzzxkzzzzgz", "5-7 h: hjmhhhhh", "10-12 l: lllllllpllllwmwwb", "2-6 j: hfjqlm", "6-9 n: mnnqfnpkw", "7-8 c: cckcqchscc", "4-5 q: qqqqq", "11-13 d: dddddddjddddgd", "3-11 q: qvqlqqqqlknnqhq", "1-4 h: hnphh", "9-15 f: ffffffffffffffffffff", "4-7 p: pxpvbpcz", "9-14 g: ggggggggwgggggg", "1-2 w: wwwwwwwwww", "4-5 g: ggcgggvg", "11-15 k: kkkgtkcjmwkrkkkkkk", "3-4 w: wzww", "8-10 k: kkkkkkkwjqtkkkkkkkk", "5-6 n: djnfkmzzbc", "12-18 x: jxxnxhqxxxxsxxxbxxxx", "13-14 p: hbmspppspppkgp", "15-16 d: dtddddddddddddpkd", "4-6 q: bvqqcpq", "2-4 h: dqhsjpp", "10-11 l: llllllklwll", "4-5 h: dhhcdh", "3-5 t: mvntxdsxftt", "5-9 w: wxpqwwqwjf", "14-15 c: cccccccccccccccc", "6-8 b: bbtbbdbwb", "4-7 g: gdpqgrj", "1-6 p: qppppppp", "1-14 s: sqssssbrszsskmss", "7-9 n: nnnnnnvsn", "4-7 r: nrltrrhndcbrr", "8-18 d: dddddddpdddddddddfdd", "1-3 b: bbbbb", "6-7 f: fpfplqf", "2-8 j: jjjjjjjsj", "3-5 h: ndhwxnb", "4-5 z: zzcztzv", "5-6 c: ccccck", "3-4 q: qjqqqqq", "4-19 q: nsqqgqpgmjsmbkpzwvl", "13-15 r: rrrrrrrrrrrrrrrrr", "6-12 k: kckvkkhkfbkkkkkkr", "3-5 x: xxxxsxx", "3-4 s: ssff", "2-7 l: pgwllllhlxwl", "8-18 z: zzzzzzzzzzzzzzzzzz", "1-11 c: fhhcjcchrlxcpcc", "3-4 w: zhcjwjnwgwg", "1-7 x: zxlpxmxt", "12-13 k: jskdrkmhkkkndk", "1-8 x: xmpwwvrx", "10-12 r: kprvkppwczjt", "1-12 b: sbbbvgbbzgbtbklsbb", "1-2 z: zzgqpwhzrhnrgjrgrstz", "9-11 s: ssssxssssss", "3-10 k: sbksdgcrkkxfgc", "7-8 n: nnnnnnzf", "13-14 j: jjjjjjjjjjjjjj", "3-4 f: fmff", "13-14 k: nkkkkkkkkkkkkk", "2-3 w: zbtww", "1-5 k: kgmkkrckdp", "4-5 p: xphlp", "3-4 k: tzzh", "2-6 s: jssvsxdsszsthspslztk", "1-3 z: mzzzqzzz", "5-7 h: hkhhwgwlchnh", "8-14 v: dvvvhggtsvdwgx", "3-8 d: dxdrkjcd", "2-4 d: pkmb", "8-11 b: bbbbbbbmbbkbbb", "3-12 j: wxljjrjjjjhzj", "5-17 c: wbnpcnngzcgcgbqwvkh", "2-6 q: sqlbhq", "1-3 n: nlbqgnhcnnznj", "2-3 k: gkncgkn", "4-6 g: gwxglggzhnq", "3-14 j: dgjcfxrjcgwjmjdmggk", "2-3 z: qvnrlzs", "11-12 k: kdtwlzkjfdkkfkn", "15-18 b: bqpxdxbbhkbnfdbfpb", "5-15 j: jjjjgjjrjjjjjmmjjjj", "8-13 t: tthtpttttmtktttt", "6-8 f: wxbpffcf", "8-9 b: bbbbbbbwr", "8-14 l: xmlrtbphkjrxrg", "15-18 x: qxxxlpxxmxxhmmxxxx", "7-11 c: cccccsclcrc", "1-4 k: kkkkkk", "3-4 j: jdqq", "11-15 b: vgfqbwsrkhgjhjv", "13-16 k: kkkkkfkkkkckkkkkkk", "18-20 b: lwwgrrbtwbcvdhbwbmbf", "7-12 m: mmmmmmmmmmmmmm", "2-12 l: pllndrxbzqzlmzstmbvw", "3-14 r: rrrrrrrrrrrrrrrrrrr", "1-16 p: ppppppppppppppppppp", "5-9 x: mtbxxxjxxxxd", "2-4 r: rrrrrrrrrr", "3-4 k: kfkz", "3-5 m: mgmmmm", "11-12 l: blcpjbfvzjll", "2-4 z: tzzzd", "5-9 z: zzzzqzzzrzjzzzzz", "10-14 m: mmmmmmmmmmmmvcmmm", "2-7 w: jwchkdws", "12-14 p: mdpppjpppppppppppr", "15-18 v: vvvvvvvvvvvvvvtvvcv", "4-10 p: pppfmbkhqppkvzpp", "3-17 p: ppppppjppppppppppppp", "3-6 r: cjvklrqrnhkk", "1-3 s: sssss", "2-6 d: dqddctdjd", "12-16 k: kkprkzdksmkbkfxd", "8-13 g: pggggggzwgggwgggmg", "13-16 b: sbwbbpbhbbbbdbbgb", "10-14 c: ccccccccccccccccc", "4-8 j: cxkxmzcrsnvbkfvmlk", "5-13 l: lmdnlnkbtspvll", "3-4 g: ggmgmmhwfmg", "5-7 d: ddddjdp", "11-15 v: vplvcvvdvhvmvtv", "10-12 r: rhrrrjrslrrnrqrrrzrn", "2-6 x: xhxxxxxx", "5-10 s: zmrkslzcdsrssfsssss", "5-7 r: rrrrrrr", "3-4 g: ggpx", "14-15 s: sssssssssssssmb", "1-5 w: smclhnwxff", "4-9 l: lllcllllqvlllll", "1-4 s: lhpd", "12-14 b: cwbdzbbbbdbfqc", "5-13 b: qtbsqbvjbcxbkznrlb", "10-12 f: fffpfffffxfqffffff", "15-16 x: xkxdqrxlwxxvxtxx", "14-15 p: ppjpgdppppppppgpp", "10-16 n: nnwnnnnnnnrqnnnn", "5-7 k: kwwzqdn", "2-3 m: mchvnq", "6-10 h: vhhxzphthw", "5-6 r: slpcrr", "17-18 b: bbbbcmbbbbbbbbbbpzbb", "3-5 s: wfscss", "9-12 r: rfjrrtrrrrkrr", "1-8 t: tstztmttgtttfvt", "8-9 f: fffffffjnff", "7-13 l: lkxclrlgldxllcll", "10-12 s: bssssdssbsxc", "8-12 t: tttttttwtttttt", "3-11 d: ndrndqqzqhdnjf", "1-10 g: xgqgggqggfmgggggggmg", "8-9 x: kttxbdnxxxcfjxxrxgxx", "4-10 n: ndjrqmcxnh", "12-13 c: txpwlvzhxhcwbpg", "8-9 z: crsqsztzz", "5-6 b: vbwbjb", "4-5 v: vvvvv", "2-6 d: tdsddddx", "4-10 r: hzqzrrnwrrhrrfrfwrr", "11-18 g: ggggggggggggggggggg", "1-7 z: zhcppxz", "7-9 r: rrwrrrrrr", "1-9 l: hlllllllqllpllllllvd", "1-4 g: jklqjgvtxsggfjggg", "2-18 m: mmqmmmjkmmmlkmmmlvmm", "1-8 w: kwgwxzshwwwc", "6-7 f: fffzfgk", "6-8 t: ttttthtcttk", "3-14 s: hsszwzdxkhsphrtrcbs", "4-10 l: lllllbllllll", "3-8 g: gggdgggkj", "1-4 x: xxxcx", "3-7 k: tqhknrkxjxqv", "11-19 p: pwwppgsztplgssgtpsp", "11-12 k: kkkkkkkkkklgkkk", "4-16 h: tnhhcgldlnzngvhb", "8-11 w: znwwwwwwwwwwww", "8-9 m: mbtmmmccjvmtm", "14-15 l: lllllllllltllfw", "11-16 z: kfslzzwszdmsnptcz", "5-6 g: qgggggk", "3-8 k: kkmdkwkd", "9-14 v: vvvvvvcvnvvvxwvfvv", "3-7 v: vvtvvvcmvvvvxvvvv", "6-7 l: dnfjfzfhhn", "9-10 g: mtzwcwgfjgshqg", "2-9 s: bsgcgjrqs", "3-4 c: ccpc", "9-14 n: nnnnnnnnnnnnnnknnnn", "4-5 p: tppxsp", "4-6 x: xmxxnxxw", "3-4 m: mmmm", "15-16 s: swshctdsksqgcsqsr", "2-4 x: xxxxxx", "6-7 h: hhvhhhh", "8-15 n: nnmvfnhnlndnwgjthmb", "4-5 z: fvzzzw", "1-12 m: wmjwgmmlmmmm", "7-16 n: xhnxjkjprznthqmdnwnl", "6-9 b: bbbbbbbbb", "17-19 d: ddwdlddkrgdkdddpqdj", "6-9 b: bbbbbbbbbb", "7-19 l: llvllllllclllflllll", "7-9 h: lhhhhhfhhhh", "7-10 m: hsdgdmmqqncdcgfxpb", "7-8 h: hqzcmvhw", "3-7 l: clkllltlwlp", "1-2 g: gsgggggqgg", "6-8 j: qjjqqcjjjjtjs", "11-14 r: nrlhdrrlbrrxgg", "3-10 s: tsxfscsfsws", "6-7 t: tstttdtbt", "6-7 r: rrcxrplkrnrrrbrpl", "3-11 b: pbwvpbkbzdbwbvlb", "5-11 d: wgwhfxtjmdfd", "7-12 m: chmmmrmrmxmqjcpmb", "1-2 n: nntnnn", "4-12 l: lllllllllllnllll", "4-5 c: ccchc"}
	result := passwordsValidNum(list)
	expected := 447

	if result != expected {
		t.Errorf("When getting the number of valid passwords, got: %v, want: %v.", result, expected)
	}

}
