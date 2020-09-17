package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

var pr = flag.String("pr", "", "The PR number to approve")

func init() {
	log.SetFlags(0)
	flag.Parse()
}

func main() {
	if *pr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	emoji := emojis[rand.Intn(len(emojis))]

	cmd := exec.Command("gh", "pr", "review", *pr, "-c", "-b", emoji)
	buf := bytes.Buffer{}
	cmd.Stderr = &buf
	err := cmd.Run()
	if err != nil {
		log.Fatal(buf.String())
	}

	cmd = exec.Command("gh", "pr", "review", *pr, "-a")
	buf = bytes.Buffer{}
	cmd.Stderr = &buf
	err = cmd.Run()
	if err != nil {
		log.Fatal(buf.String())
	}
	log.Printf("PR %s approved with %s", *pr, emoji)
}

var emojis = []string{
	"😈",
	"😎",
	"💫",
	"👻",
	"👹",
	"🙉",
	"👿",
	"❤️",
	"💩",
	"👺",
	"💛",
	"💕",
	"😬",
	"😏",
	"😊",
	"💯",
	"💥",
	"👾",
	"💣",
	"💟",
	"😂",
	"😍",
	"🤔",
	"😘",
	"🙄",
	"😀",
	"😁",
	"🤣",
	"😃",
	"😄",
	"😅",
	"😋",
	"😗",
	"😆",
	"😉",
	"😙",
	"😚",
	"☺️",
	"🙂",
	"🤗",
	"😐",
	"😑",
	"😶",
	"😣",
	"😥",
	"😮",
	"🤐",
	"😯",
	"😪",
	"😫",
	"😴",
	"😌",
	"😛",
	"😜",
	"😝",
	"🤤",
	"😒",
	"😓",
	"😔",
	"😕",
	"🙃",
	"☹️",
	"🤑",
	"😲",
	"🙁",
	"😖",
	"😞",
	"😟",
	"😤",
	"😢",
	"😭",
	"😦",
	"😧",
	"😨",
	"😩",
	"😰",
	"😱",
	"😳",
	"😵",
	"😡",
	"😠",
	"😷",
	"🤒",
	"🤢",
	"🤕",
	"🤧",
	"😇",
	"🤠",
	"🤡",
	"🤥",
	"🤓",
	"💀",
	"👽",
	"🤖",
	"😺",
	"😸",
	"😹",
	"😻",
	"😼",
	"😽",
	"🙀",
	"😿",
	"😾",
	"💨",
	"🕳",
	"💋",
	"💌",
	"🙈",
	"☠️",
	"🙊",
	"💦",
	"👁️‍🗨️",
	"💓",
	"💔",
	"💖",
	"💗",
	"💘",
	"💙",
	"💚",
	"💜",
	"💝",
	"💞",
	"💢",
	"💤",
	"💬",
	"💭",
	"🖤",
	"🗨",
	"🗯",
	"🧡",
	"❣️",
	"🤪",
	"🤬",
	"🤯",
	"🥵",
	"🥶",
	"🤨",
	"🤩",
	"🥰",
	"🤫",
	"🤭",
	"🥳",
	"🥴",
	"🥺",
	"🧐",
	"🤮",
	"🥱",
	"🤎",
	"🤍",
	"👀",
	"🕵",
	"👆",
	"✌",
	"👶",
	"💪",
	"👱",
	"👏",
	"👊",
	"🙏",
	"🚣",
	"👸",
	"🤷",
	"👤",
	"👦",
	"👧",
	"👨",
	"👩",
	"👴",
	"👵",
	"👨‍⚕️",
	"👩‍⚕️",
	"👨‍🎓",
	"👩‍🎓",
	"👨‍⚖️",
	"👩‍⚖️",
	"👨‍🌾",
	"👩‍🌾",
	"👨‍🍳",
	"👩‍🍳",
	"👨‍🔧",
	"👩‍🔧",
	"👨‍🏭",
	"👩‍🏭",
	"👨‍💼",
	"👩‍💼",
	"👨‍🔬",
	"👩‍🔬",
	"👨‍💻",
	"👩‍💻",
	"👨‍🎤",
	"👩‍🎤",
	"👨‍🎨",
	"👩‍🎨",
	"👨‍✈️",
	"👩‍✈️",
	"👨‍🚀",
	"👩‍🚀",
	"👨‍🚒",
	"👩‍🚒",
	"👮",
	"👮‍♂️",
	"👮‍♀️",
	"🕵️‍♀️",
	"💂",
	"💂‍♂️",
	"💂‍♀️",
	"👷",
	"👷‍♂️",
	"👷‍♀️",
	"🤴",
	"👳",
	"👳‍♀️",
	"👲",
	"👱‍♀️",
	"🗣",
	"💇‍♀️",
	"🤞",
	"👂",
	"👁",
	"👪",
	"👣",
	"🤝",
	"🏇",
	"🤛",
	"🏌️‍♂️",
	"👄",
	"💅",
	"👃",
	"👌",
	"👐",
	"🏌",
	"🛌",
	"🛀",
	"🤚",
	"✊",
	"✋",
	"🙌",
	"🤜",
	"🤘",
	"⛷",
	"🏂",
	"👎",
	"👍",
	"👅",
	"🖖",
	"👋",
	"🏌️‍♀️",
	"✍",
	"👼",
	"👇",
	"👈",
	"👉",
	"👰",
	"👥",
	"💑",
	"👨‍❤️‍👨",
	"👩‍❤️‍👩",
	"👨‍👦",
	"👨‍👦‍👦",
	"👨‍👧",
	"👨‍👧‍👦",
	"👨‍👧‍👧",
	"👨‍👨‍👦",
	"👨‍👨‍👦‍👦",
	"👨‍👨‍👧",
	"👨‍👨‍👧‍👦",
	"👨‍👨‍👧‍👧",
	"👨‍👩‍👦",
	"👨‍👩‍👦‍👦",
	"👨‍👩‍👧",
	"👨‍👩‍👧‍👦",
	"👩‍👦",
	"👩‍👦‍👦",
	"👩‍👧",
	"👩‍👧‍👦",
	"👩‍👧‍👧",
	"👩‍👩‍👦",
	"👩‍👩‍👦‍👦",
	"👩‍👩‍👧",
	"👩‍👩‍👧‍👦",
	"☝",
	"💏",
	"👨‍❤️‍💋‍👨",
	"👩‍❤️‍💋‍👩",
	"👫",
	"🙇‍♂️",
	"🕺",
	"🤦‍♂️",
	"🙍‍♂️",
	"🙅‍♂️",
	"🙆‍♂️",
	"💇‍♂️",
	"💆‍♂️",
	"🕴",
	"🤵",
	"🙎‍♂️",
	"🙋‍♂️",
	"🏃‍♂️",
	"🤷‍♂️",
	"💁‍♂️",
	"🚶‍♂️",
	"🖕",
	"🤶",
	"👯",
	"🙇",
	"🤦",
	"🙍",
	"🙅",
	"🙆",
	"💇",
	"💆",
	"🙎",
	"🙋",
	"🏃",
	"💁",
	"🚶",
	"🤰",
	"🖐",
	"🎅",
	"🤳",
	"👬",
	"👭",
	"🙇‍♀️",
	"💃",
	"🤦‍♀️",
	"🙍‍♀️",
	"🙅‍♀️",
	"🙆‍♀️",
	"🤙",
	"🤟",
	"🤲",
	"🦴",
	"🦵",
	"🦶",
	"🦷",
	"🧠",
	"👨‍👩‍👧‍👧",
	"👩‍👩‍👧‍👧",
	"👩‍❤️‍👨",
	"👩‍❤️‍💋‍👨",
	"🧑",
	"🧒",
	"🧓",
	"🏃‍♀️",
	"👯‍♀️",
	"👯‍♂️",
	"💆‍♀️",
	"🚶‍♀️",
	"🧖",
	"🧖‍♀️",
	"🧖‍♂️",
	"🧗",
	"🧗‍♀️",
	"🧗‍♂️",
	"🧘",
	"🧘‍♀️",
	"🧘‍♂️",
	"🦸",
	"🦸‍♀️",
	"🦸‍♂️",
	"🦹",
	"🦹‍♀️",
	"🦹‍♂️",
	"🧙",
	"🧙‍♀️",
	"🧙‍♂️",
	"🧚",
	"🧚‍♀️",
	"🧚‍♂️",
	"🧛",
	"🧛‍♀️",
	"🧛‍♂️",
	"🧜",
	"🧜‍♀️",
	"🧜‍♂️",
	"🧝",
	"🧝‍♀️",
	"🧝‍♂️",
	"🧞",
	"🧞‍♀️",
	"🧞‍♂️",
	"🧟",
	"🧟‍♀️",
	"🧟‍♂️",
	"💁‍♀️",
	"🙋‍♀️",
	"🙎‍♀️",
	"🤷‍♀️",
	"👨‍🏫",
	"👨‍🦰",
	"👨‍🦱",
	"👨‍🦲",
	"👨‍🦳",
	"👩‍🏫",
	"👩‍🦰",
	"👩‍🦱",
	"👩‍🦲",
	"👩‍🦳",
	"👱‍♂️",
	"👳‍♂️",
	"🕵️‍♂️",
	"🤱",
	"🧔",
	"🧕",
	"🏄",
	"🏄‍♀️",
	"🏄‍♂️",
	"🏊",
	"🏊‍♀️",
	"🏊‍♂️",
	"🏋",
	"🏋️‍♀️",
	"🏋️‍♂️",
	"🚣‍♀️",
	"🚣‍♂️",
	"🚴",
	"🚴‍♀️",
	"🚴‍♂️",
	"🚵",
	"🚵‍♀️",
	"🚵‍♂️",
	"🤸",
	"🤸‍♀️",
	"🤸‍♂️",
	"🤹",
	"🤹‍♀️",
	"🤹‍♂️",
	"🤺",
	"🤼",
	"🤼‍♀️",
	"🤼‍♂️",
	"🤽",
	"🤽‍♀️",
	"🤽‍♂️",
	"🤾",
	"🤾‍♀️",
	"🤾‍♂️",
	"⛹",
	"⛹️‍♀️",
	"⛹️‍♂️",
	"🤏",
	"🦾",
	"🦿",
	"🦻",
	"🧏",
	"🧏‍♂️",
	"🧏‍♀️",
	"🧍",
	"🧍‍♂️",
	"🧍‍♀️",
	"🧎",
	"🧎‍♂️",
	"🧎‍♀️",
	"👨‍🦯",
	"👩‍🦯",
	"👨‍🦼",
	"👩‍🦼",
	"👨‍🦽",
	"👩‍🦽",
	"🧑‍🤝‍🧑",
	"🧑‍🦰",
	"🧑‍🦱",
	"🧑‍🦳",
	"🧑‍🦲",
	"🧑‍⚕️",
	"🧑‍🎓",
	"🧑‍🏫",
	"🧑‍⚖️",
	"🧑‍🌾",
	"🧑‍🍳",
	"🧑‍🔧",
	"🧑‍🏭",
	"🧑‍💼",
	"🧑‍🔬",
	"🧑‍💻",
	"🧑‍🎤",
	"🧑‍🎨",
	"🧑‍✈️",
	"🧑‍🚀",
	"🧑‍🚒",
	"🧑‍🦯",
	"🧑‍🦼",
	"🧑‍🦽",
	"🐼",
	"🌵",
	"🐨",
	"🐀",
	"🌹",
	"🐅",
	"🌷",
	"🐜",
	"🐤",
	"🦇",
	"🐻",
	"🐦",
	"🌼",
	"🐡",
	"🐗",
	"💐",
	"🐛",
	"🦋",
	"🐪",
	"🐈",
	"🐱",
	"🌸",
	"🐔",
	"🐿",
	"🐄",
	"🐮",
	"🐊",
	"🌳",
	"🐕",
	"🐶",
	"🐬",
	"🕊",
	"🐉",
	"🐲",
	"🦆",
	"🦅",
	"🐘",
	"🌲",
	"🐑",
	"🍂",
	"🐟",
	"🍀",
	"🦊",
	"🐸",
	"🐥",
	"🐐",
	"🦍",
	"🐹",
	"🐣",
	"🌿",
	"🌺",
	"🐝",
	"🐎",
	"🐴",
	"🐞",
	"🍃",
	"🐆",
	"🦁",
	"🦎",
	"🍁",
	"🐒",
	"🐵",
	"🐁",
	"🐭",
	"🐙",
	"🦉",
	"🐂",
	"🌴",
	"🐾",
	"🐧",
	"🐖",
	"🐷",
	"🐽",
	"🐩",
	"🐇",
	"🐰",
	"🐏",
	"🦏",
	"🐓",
	"🏵",
	"🦂",
	"🌱",
	"☘️",
	"🌾",
	"🐌",
	"🐍",
	"🕷",
	"🕸",
	"🐚",
	"🐳",
	"🌻",
	"🐯",
	"🐠",
	"🦃",
	"🐢",
	"🐫",
	"🦄",
	"🐃",
	"🐋",
	"💮",
	"🥀",
	"🐺",
	"🦚",
	"🦜",
	"🦢",
	"🦗",
	"🦟",
	"🦠",
	"🦌",
	"🦒",
	"🦓",
	"🦔",
	"🦘",
	"🦙",
	"🦛",
	"🦝",
	"🦡",
	"🦈",
	"🦕",
	"🦖",
	"🦧",
	"🦮",
	"🐕‍🦺",
	"🦥",
	"🦦",
	"🦨",
	"🦩",
	"🍭",
	"🍏",
	"🎂",
	"🥑",
	"🍳",
	"🍌",
	"🍞",
	"🍰",
	"🍒",
	"🍓",
	"🍐",
	"🍕",
	"🍋",
	"🌰",
	"🏺",
	"🍼",
	"🥓",
	"🥖",
	"🍺",
	"🍱",
	"🍾",
	"🌯",
	"🍬",
	"🥕",
	"🧀",
	"🍫",
	"🍻",
	"🥂",
	"🍸",
	"🍚",
	"🍪",
	"🦀",
	"🥐",
	"🥒",
	"🍛",
	"🍮",
	"🍡",
	"🍩",
	"🌽",
	"🍆",
	"🍥",
	"🍴",
	"🍽",
	"🍟",
	"🍤",
	"🥛",
	"🍇",
	"🥗",
	"🍔",
	"🍯",
	"☕",
	"🌭",
	"🌶",
	"🍨",
	"🔪",
	"🥝",
	"🍖",
	"🍈",
	"🍄",
	"🍢",
	"🥞",
	"🍑",
	"🥜",
	"🍿",
	"🍲",
	"🥔",
	"🍗",
	"🍎",
	"🍙",
	"🍘",
	"🍠",
	"🍶",
	"🍧",
	"🦐",
	"🍦",
	"🍝",
	"🥄",
	"🦑",
	"🍜",
	"🍣",
	"🌮",
	"🍊",
	"🍵",
	"🍅",
	"🍹",
	"🥃",
	"🍉",
	"🍷",
	"🦞",
	"🥢",
	"🥤",
	"🥟",
	"🥠",
	"🥡",
	"🥮",
	"🍍",
	"🥥",
	"🥭",
	"🥘",
	"🥙",
	"🥚",
	"🥣",
	"🥨",
	"🥩",
	"🥪",
	"🥫",
	"🥯",
	"🧂",
	"🥧",
	"🧁",
	"🥦",
	"🥬",
	"🧄",
	"🧅",
	"🧇",
	"🧆",
	"🧈",
	"🦪",
	"🧃",
	"🧉",
	"🧊",
	"🔵",
	"⚪",
	"♈",
	"♉",
	"♊",
	"♋",
	"♌",
	"♍",
	"♎",
	"♏",
	"♐",
	"♑",
	"♒",
	"♓",
	"🔂",
	"🔁",
	"🔄",
	"🔠",
	"⚜️",
	"⬆️",
	"🛄",
	"🛃",
	"🛅",
	"🛂",
	"🅰️",
	"🅱️",
	"🅾️",
	"🅿️",
	"🆎",
	"🆑",
	"🆒",
	"🆓",
	"🆔",
	"🆕",
	"🆖",
	"🆗",
	"🆘",
	"🆙",
	"🆚",
	"🈁",
	"🈂️",
	"🈚",
	"🈯",
	"🈲",
	"🈳",
	"🈴",
	"🈵",
	"🈶",
	"🈷️",
	"🈸",
	"🈹",
	"🈺",
	"🉐",
	"🉑",
	"🔡",
	"🔢",
	"🔣",
	"🔤",
	"ℹ️",
	"Ⓜ️",
	"㊗️",
	"㊙️",
	"🔃",
	"🔙",
	"🔚",
	"🔛",
	"🔜",
	"🔝",
	"↔",
	"↕",
	"↖️",
	"↗️",
	"↘️",
	"↙️",
	"↩️",
	"↪️",
	"➡️",
	"⤴️",
	"⤵️",
	"⬅️",
	"⬇️",
	"🎦",
	"📳",
	"📴",
	"📶",
	"🔀",
	"🔅",
	"🔆",
	"🔼",
	"🔽",
	"⏏️",
	"⏩",
	"⏪",
	"⏫",
	"⏬",
	"⏭",
	"⏮",
	"⏯",
	"⏸",
	"⏹",
	"⏺",
	"▶️",
	"◀️",
	"💠",
	"🔘",
	"🔲",
	"🔳",
	"🔴",
	"🔶",
	"🔷",
	"🔸",
	"🔹",
	"🔺",
	"🔻",
	"▪",
	"▫",
	"◻️",
	"◼️",
	"◽",
	"◾",
	"⚫",
	"⬛",
	"⬜",
	"#️⃣",
	"*️⃣",
	"0️⃣",
	"1️⃣",
	"2️⃣",
	"3️⃣",
	"4️⃣",
	"5️⃣",
	"6️⃣",
	"7️⃣",
	"8️⃣",
	"9️⃣",
	"🔟",
	"💱",
	"💲",
	"📛",
	"🔰",
	"🔱",
	"‼",
	"⁉️",
	"™",
	"☑️",
	"♻️",
	"♾️",
	"⚕️",
	"✅",
	"✔️",
	"✖️",
	"✳️",
	"✴️",
	"❇️",
	"❌",
	"❎",
	"❓",
	"❔",
	"❕",
	"❗",
	"➕",
	"➖",
	"➗",
	"➰",
	"➿",
	"⭕",
	"〰️",
	"〽️",
	"🔯",
	"🕉",
	"🕎",
	"🛐",
	"☦️",
	"☪️",
	"☮️",
	"☯️",
	"☸️",
	"⚛️",
	"✝️",
	"✡️",
	"🏧",
	"🚮",
	"🚰",
	"🚹",
	"🚺",
	"🚻",
	"🚼",
	"🚾",
	"♿",
	"📵",
	"🔞",
	"🚫",
	"🚭",
	"🚯",
	"🚱",
	"🚳",
	"🚷",
	"🚸",
	"☢️",
	"☣️",
	"⚠️",
	"⛔",
	"⛎",
	"🟠",
	"🟡",
	"🟢",
	"🟣",
	"🟤",
	"🟥",
	"🟧",
	"🟨",
	"🟩",
	"🟦",
	"🟪",
	"🟫",
	"🌎",
	"🌟",
	"🚀",
	"🌞",
	"❄️",
	"🌈",
	"🌜",
	"🔥",
	"⌛",
	"🚘",
	"🌙",
	"⭐",
	"🚨",
	"🚡",
	"✈️",
	"🛬",
	"🛫",
	"⏰",
	"🚑",
	"⚓",
	"🚛",
	"🚗",
	"🏦",
	"💈",
	"🏖",
	"🛎",
	"🚲",
	"🌉",
	"🏗",
	"🚌",
	"🚏",
	"🏕",
	"🎠",
	"🏰",
	"⛪",
	"🏙",
	"🌆",
	"🏛",
	"🌂",
	"☁️",
	"🌩",
	"⛈",
	"🌧",
	"🌨",
	"☄️",
	"🚧",
	"🏪",
	"🚚",
	"🏬",
	"🏚",
	"🏜",
	"🏝",
	"💧",
	"🏭",
	"🎡",
	"⛴",
	"🚒",
	"🌓",
	"🌛",
	"🌫",
	"🌁",
	"⛲",
	"⛽",
	"🌕",
	"🌝",
	"🌏",
	"🌍",
	"🌐",
	"🚁",
	"⚡",
	"🚄",
	"🚅",
	"🚥",
	"🏥",
	"🏨",
	"⏳",
	"🏘",
	"🏠",
	"🏡",
	"🏯",
	"🏣",
	"🕋",
	"🛴",
	"🌗",
	"🚈",
	"🚂",
	"🏩",
	"🕰",
	"🗾",
	"🚇",
	"🌌",
	"🚐",
	"🚝",
	"🕌",
	"🛥",
	"🛵",
	"🏍",
	"🗻",
	"⛰",
	"🚠",
	"🚞",
	"🏞",
	"🌑",
	"🌚",
	"🌃",
	"🏢",
	"🚍",
	"🚔",
	"🚖",
	"🛳",
	"🚓",
	"🏤",
	"🏎",
	"🚃",
	"🛤",
	"🎢",
	"⛵",
	"🛰",
	"🏫",
	"💺",
	"⛩",
	"🚢",
	"🌠",
	"🛩",
	"🏔",
	"☃️",
	"⛄",
	"🚤",
	"🏟",
	"🚉",
	"🗽",
	"⏱",
	"☀️",
	"⛅",
	"🌥",
	"🌦",
	"🌤",
	"🌅",
	"🌄",
	"🌇",
	"🚟",
	"🕍",
	"🚕",
	"⛺",
	"🌡",
	"⏲",
	"🗼",
	"🌪",
	"🚜",
	"🚆",
	"🚊",
	"🚋",
	"🚎",
	"☂️",
	"⛱",
	"☔",
	"🚦",
	"🌋",
	"🌘",
	"🌖",
	"⌚",
	"🌊",
	"🌒",
	"🌔",
	"💒",
	"🌬",
	"🗺",
	"🧳",
	"🧱",
	"🧭",
	"🎪",
	"♨️",
	"🌀",
	"🕐",
	"🕑",
	"🕒",
	"🕓",
	"🕔",
	"🕕",
	"🕖",
	"🕗",
	"🕘",
	"🕙",
	"🕚",
	"🕛",
	"🕜",
	"🕝",
	"🕞",
	"🕟",
	"🕠",
	"🕡",
	"🕢",
	"🕣",
	"🕤",
	"🕥",
	"🕦",
	"🕧",
	"🛸",
	"🚙",
	"🛑",
	"🛢",
	"🛣",
	"🛹",
	"🛶",
	"🛕",
	"🦽",
	"🦼",
	"🛺",
	"🪂",
	"🪐",
	"💽",
	"🔗",
	"☎️",
	"📧",
	"📼",
	"🎥",
	"📹",
	"🔦",
	"🔎",
	"🔬",
	"📻",
	"🎬",
	"📺",
	"📠",
	"👑",
	"💵",
	"📀",
	"🎧",
	"👠",
	"🔑",
	"💄",
	"💊",
	"🗳",
	"📊",
	"🛁",
	"🔋",
	"🛏",
	"👙",
	"✒️",
	"📘",
	"🔖",
	"📑",
	"📚",
	"💼",
	"📅",
	"📷",
	"📸",
	"🕯",
	"🗃",
	"📇",
	"🗂",
	"📉",
	"📈",
	"📋",
	"📕",
	"📪",
	"📫",
	"👝",
	"🖱",
	"🎛",
	"🛋",
	"🖍",
	"💳",
	"🖥",
	"🚪",
	"👗",
	"🔌",
	"✉️",
	"📩",
	"💶",
	"🗄",
	"📁",
	"🎞",
	"📽",
	"💾",
	"🖋",
	"💎",
	"👓",
	"🎓",
	"📗",
	"🔨",
	"⚒️",
	"👜",
	"📥",
	"📨",
	"👖",
	"⌨️",
	"👘",
	"🏷",
	"💻",
	"🔍",
	"🎚",
	"💡",
	"🖇",
	"🔒",
	"🔐",
	"🔏",
	"👞",
	"📝",
	"🗿",
	"📱",
	"📲",
	"💰",
	"💸",
	"👔",
	"📰",
	"📓",
	"📔",
	"🗝",
	"📖",
	"📂",
	"📭",
	"📬",
	"💿",
	"📙",
	"📤",
	"📦",
	"📄",
	"📃",
	"📟",
	"🖌",
	"📎",
	"🖊",
	"✏️",
	"⛏",
	"📯",
	"📮",
	"💷",
	"📿",
	"🖨",
	"👛",
	"📌",
	"🏮",
	"⛑",
	"💍",
	"🗞",
	"📍",
	"👟",
	"📡",
	"🎒",
	"✂️",
	"📜",
	"🛍",
	"🚿",
	"🗓",
	"🗒",
	"📏",
	"🎙",
	"🕶",
	"👕",
	"📆",
	"📞",
	"🔭",
	"🚽",
	"🎩",
	"🖲",
	"📐",
	"🔓",
	"🗑",
	"👢",
	"👚",
	"👒",
	"👡",
	"💴",
	"📒",
	"🥼",
	"🥽",
	"🥾",
	"🥿",
	"🧢",
	"🧣",
	"🧤",
	"🧥",
	"🧦",
	"🧮",
	"🛒",
	"🧯",
	"🧴",
	"🧷",
	"🧹",
	"🧺",
	"🧻",
	"🧼",
	"🧽",
	"💉",
	"💹",
	"🧾",
	"🎤",
	"🎵",
	"🎶",
	"🎼",
	"🎷",
	"🎸",
	"🎹",
	"🎺",
	"🎻",
	"🥁",
	"🚬",
	"⚰️",
	"⚱️",
	"🧪",
	"🧫",
	"🧬",
	"⚗️",
	"📢",
	"📣",
	"🔇",
	"🔈",
	"🔉",
	"🔊",
	"🔔",
	"🔕",
	"🏹",
	"🔧",
	"🔩",
	"🔫",
	"🗜",
	"🗡",
	"🛠",
	"🛡",
	"🧰",
	"🧲",
	"⚔️",
	"⚖️",
	"⚙️",
	"⛓",
	"🦺",
	"🥻",
	"🩱",
	"🩲",
	"🩳",
	"🩰",
	"🪕",
	"🪔",
	"🪓",
	"🦯",
	"🩸",
	"🩹",
	"🩺",
	"🪑",
	"🪒",
	"✨",
	"🎆",
	"🎨",
	"🎈",
	"⚾",
	"🎲",
	"🎀",
	"🎏",
	"🎄",
	"🎊",
	"🖼",
	"🎎",
	"🕹",
	"🎑",
	"🎉",
	"🎍",
	"🎇",
	"🎋",
	"🎐",
	"🎁",
	"🎭",
	"🧵",
	"🧶",
	"🎖",
	"🏅",
	"🏆",
	"🥇",
	"🥈",
	"🥉",
	"🎃",
	"🎗",
	"🎟",
	"🎫",
	"🧧",
	"🧨",
	"🀄",
	"🃏",
	"🎮",
	"🎯",
	"🎰",
	"🎱",
	"🎴",
	"🔮",
	"🧩",
	"🧸",
	"🧿",
	"♟️",
	"♠️",
	"♣️",
	"♥️",
	"♦️",
	"🎣",
	"🎳",
	"🎽",
	"🎾",
	"🎿",
	"🏀",
	"🏈",
	"🏉",
	"🏏",
	"🏐",
	"🏑",
	"🏒",
	"🏓",
	"🏸",
	"🛷",
	"🥅",
	"🥊",
	"🥋",
	"🥌",
	"🥍",
	"🥎",
	"🥏",
	"⚽",
	"⛳",
	"⛸",
	"🤿",
	"🪀",
	"🪁",
	"🏁",
	"🎌",
	"🏳",
	"🏳️‍🌈",
	"🏴",
	"🏴‍☠️",
	"🚩",
	"🏴󠁧󠁢󠁥󠁮󠁧󠁿",
	"🏴󠁧󠁢󠁳󠁣󠁴󠁿",
	"🏴󠁧󠁢󠁷󠁬󠁳󠁿",
	"🦰",
	"🦱",
	"🦲",
	"🦳",
}
