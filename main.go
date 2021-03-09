package main

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func init() {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("PR number or URL required")
	}

	pr := os.Args[1]

	emoji := emojis[rand.Intn(len(emojis))]

	var body string
	if len(os.Args) > 2 {
		body = " " + strings.Join(os.Args[2:], " ")
	}

	cmd := exec.Command("gh", "pr", "review", pr, "--body", emoji+body, "--approve")

	var buf bytes.Buffer
	cmd.Stderr = &buf
	err := cmd.Run()
	if err != nil {
		log.Fatal(buf.String())
	}

	if len(body) > 0 {
		log.Printf("PR %s approved with '%s%s'", pr, emoji, body)
	} else {
		log.Printf("PR %s approved with %s", pr, emoji)
	}
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
