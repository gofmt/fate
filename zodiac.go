package fate

import (
	"fmt"
	"strings"

	"github.com/godcong/chronos"
)

// ZodiacPig ...
const (
	ZodiacRat     = "鼠"
	ZodiacCow     = "牛"
	ZodiacTiger   = "虎"
	ZodiacRabbit  = "兔"
	ZodiacDragon  = "龙"
	ZodiacSnake   = "蛇"
	ZodiacHorse   = "马"
	ZodiacSheep   = "羊"
	ZodiacMonkey  = "猴"
	ZodiacChicken = "鸡"
	ZodiacDog     = "狗"
	ZodiacPig     = "猪"
)

//Zodiac 生肖
type Zodiac struct {
	Name      string
	Xi        string //喜
	XiRadical string
	Ji        string //忌
	JiRadical string
}

var zodiacList = map[string]Zodiac{
	ZodiacRat: {
		Name: ZodiacRat,
		Ji:   `丁丛东为义书二亚亞亭亮人亿仁仃仅仇今介仍从仑仕付仙仟代以仪仰仲仳仵件价任份企伂伊伍休优伙会伟传伦伫伯估伴伶伸伺似但佇佈位低住佐佑何余佛作佟佣佩佯佰佳佶佼佾使侄來例侍侑侖供依侠侦侨侬侯係促俊俋俐俗保俞俟俠信俩俭修俱俸俾倆倍倏倖候倚倡倢倨倩倪倫值偃偅假偈偉偌做停健偩偮偵偶偿傅傌傑備傢傧储傭傲傳僅僇像僑僧僵價僾儀儂億儆儉儌儐儒償優儲光免兔全养冕冤冯冻凍刘劉勉勗化午卖南卧卯印卿叢台合吗味命哇唱善喜嗎嚥在圭地场址均坊坍坎块坚坛坜坝坞坠坡坦坪坵垄型垒垠垢垣垦垫埂埃城埒埔埕域埠執基堅堆堡堤堪堯堰報場堵塊塌塑塔塗塘塚塢填境墅墉墊墜增墟墡墩墾壁壅壇壎壑壕壘壟壢壤壩士壮壯壽备复夭央妃妈妹姜姬姿娅娩婚婭媄媽嬉孔孝宋宣宰寺寿封尧岱嶪巫差巴平幸庠张張彻徉律徐得從復微徵德徹志恙悠悻惊意慢执扬抑报拒挽提揚攸教日旦旨早旬旭时旺旻昀昂昆昇昊昌明昏易昙星映春昧昭是昱昶晁時晃晉晋晏晓晖晚晟晤普景晰晴晶智暄暉暖暗暮曆曇曉曙曜曝曦書曹曼曾會未朴杜来杨杩杰東杵柏柠柯柳柴样桓棚楊楠楼榪榴槽樓樣檸次欢欤欧欲欷欸款歇歌歐歟歡氧汀沫洋涂渼漫漾潛潜澲火灿炅炎炫炬炯炳炷為炽烈烊烨烯焉焕焯焰然煜煥照熊熙熹熾燁燕營燦燨牺犧珋珠瓷留监監祀祥禡竹笃篤糕紅綻繒繕繨红绽缮缯署罵羊美羔羕羚羟羡羢群羥羨義羬羯羱羲羶羹翔耀耵耿聊肖肯胡脩腾膳臣臥臺苎苣苧英茆茉荷莘莫菟营蓦蓨薑薘蝴行袁袂袋訂許詳諮議订议许详谘貸賣贷赤赫距路輓輝辉辛辜辞辭达迈迎运进远连迪选造逢連進逸運道達遠選邁邓那邦邬邱邹邺郁郎郑郝都鄒鄔鄧鄭鄯鄴酊酒醒釘鉀鉚鋅鍚鎂鎦鑑钉钾铆锌镁镏阳陈陳陽集鞑韃頂題顶题養香馬馮馳馴駁駐駒駕駛駟駱駿騁騎騏騫騮騰驀驅驊驕驗驚驛驞驟驥驪马驯驰驱驳驶驷驹驻驾驿骂骄骅骆骊骋验骏骐骑骝骞骤骥魯鮭鮮鲁鲑鲜鵬鹏`,
	},
	ZodiacCow: {
		Name: ZodiacCow,
		Ji:   `丁丙业丛丝丬主义习书些亭人亿什仁仃仅仇今介仑仕付仙仟代令以仪仰仲仳仵件价任份企伂伊伍休优伙会伟传伦伫伯估伴伶伸伺似但佇佈佌位低住佐佑佔何余佛作佟佣佩佯佰佳佶佼佾使侄來例侍侑侖供依侦侨侬侯係促俊俋俐俗保俟信俩俪俭修俱俸俾倆倍倏倖候倚倡倢倧倨倩倪倫值偃偅假偈偉偌做停健偩偮偵偶偿傅傌傑備傢傧储傭傲傳僅僇像僑僧僵價僾儀儂億儆儉儌儐儒償優儲儷光全养冈冯准刀刊刑划列刘则刚创初判別别到制刷券刻剀剂則剋前剑剛剧副剴創劃劇劉劍劑劝功加务动励劳势勇動勗勘務勝勞勢勤勵勸匕化午协卓協南占卧县叢古召叱合吗君吟呀味命唬唱善嗎嗣嚥国國场坎坏堤報場增墡壞备夏大天夭央奂奈奎奐奖妈妹姜姝娥婚媄媽嬉它宇宗审宣宰寥審将將岡岱岳峰峻崇崙崟嵘嶪嶸巫差希干年幸庄庇应庠廖式弗弟弥弦弭弯彌彎形彤彥彦彩彪彫彬彭彰影徉御德心必忆忌忍志忙忠快念忻忽怀态怎怔怕怖怜思怡急性怨怪总恃恆恋恐恕恙恢恣恩恬恰恳悄悅悉悍悔悖悟悠悦悬悲悸悻悽情惊惜惟惠惧惭惯想惶愁愃愈愉愍意愚愛感愠愤愧慄慇慈態慌慍慎慚慢慣慧慨慰慷慾憎憐憤憩憶憾懂懇應懋懷懸懼戀戈戎戏戕或战戡截戰戲戴扆扎托扬报挂指捺掛採提揚摎支攸教斌斤断斯斷旋日旦旨早旭时旺旻昀昂昆昇昊昌明昏易昕昙星映春昧昭是昱昶晁時晃晉晋晏晒晓晖晚晟晤晨普景晰晴晶智暄暉暖暗暢暮曆曇曉曙曜曝曦曬書曹曼曾會月有朋服朔朗望朝期朧未末朱杆杉来杨杩杰杵枝架柏柔柠柰栈栩样格桨桿梧棉棕棧棹椅楊楠業榪槳槽樣樱檸櫻欢欣欤欧欲欷欸款歇歌歐歟歡此比氧汀沫洋洧浅淺渼湾準漫漻漾潛潜澎澲灣炅炘炽烊烯照熾燕燨爱爿牆片版牋牌牍牒牘牙牺犧状狀獎玉王玕玖玛玟玡玥玨玫玮环现玲玳玺玼珀珊珍珏珑珖珠班珮珺現琁球琅理琉琏琐琥琦琨琪琮琯琳琴琼瑋瑙瑚瑛瑜瑞瑟瑣瑤瑩瑪瑰瑶璀璃璇璉璋璎璞璧璨環璽璿瓊瓏瓔甯畅番疆直矇矛矜矞矢知矩矫短矯碧示礼礽社祀祁祅祇祈祉祐祕祖祜祝神祟祠祥祭祯祷祸祺祿禁禄禅禍禎福禡禦禧禪禮禱禳秘穿竹笃篤米粽糕系糾紀紂約紅紆紇紈紉紋納紐紓純紗紘紙級紛紜素紡索紧紫紮細紳紵紹紼紾終絃組結絕絜絡絢給絨絪絮統絲絹絻綏經綖綜綠綢綦綪綬維綰綱網綴綵綸綺綻綽綾綿緁緆緇緊緋緒緗緘緞締緣編緩緬緯緲練緹緻縈縉縑縕縝縣縱縴縵縷總績繆繒織繕繚繞繡繨繩繪繫繯繳繹繼繽繾續纏纓纖纘纜纠纡红纣纤纥约级纨纪纫纬纭纯纱纲纳纵纶纷纸纹纺纽纾练组绅细织终绋绍绎经绒结绕绘给绚络绝统绢绣绥继绩绪绫续绮绯绰绳维绵绶绸综绽绾绿缀缁缃缄缅缆缇缈缎缓缔缕编缘缙缜缠缣缤缦缨缪缭缮缯缱缳缴缵网罕署罵羊美羔羕羚羟羡羢群羥羨義羬羯羱羲羶羹羿翁翎習翔翕翩翼耀耵耿聆肯育胖胜胞胡胤胥胧胭胶能脈脉脩腾膠膳臆臥臧芯苎苓苧茅茉莊莘莫莹萦蒋蓦蓨蔣蕙蕥薑薘藏蟉行衣补表衫衬衽衾衿袂袋袍袒袖袗被裁裂装裆裕裘裙補裝裟裤裱裴裸製複褊褋褌褐褓褙褚褛褪褫褲褸襁襟襠襯襴視视訂託許診試詳識議订议许识诊试详責貸費質责质贷费赤赫赶趕路躬軫輝轸辉辛辜辞辭达達邪邺郁郝郡鄢鄯鄴酊醒采釘鉀鋅錼鍚鎂鑫钉钾锌镁阳陀际陇陈陳陽際隴集雉零青鞑韃頂須題顏顶须题颜養馬馮馳馴駁駐駒駕駛駟駱駿騁騎騏騫騰驀驅驊驕驗驚驛驞驟驤驥驪马驯驰驱驳驶驷驹驻驾驿骂骄骅骆骊骋验骏骐骑骞骤骥骧魠魯鮮鲁鲜鷚鷸鹨鹬齒齡齿龄`,
	},
	ZodiacTiger: {
		Name: ZodiacTiger,
		Ji:   `一万三严丰临之乏乔乡书二亚亞亨享亭亮人亿仁仅仇今介仍从仑仓仕仗付仙仟代以仪仰仲仳仵件价任份企伂伊伍休优伙会伟传伦伯估伴伶伸伺似佃但佈位低住佐佑何余佛作佟佣佩佰佳佶佼佾使侄來例侍侏侑侖侗供依侠侣侦侨侬侯侶係促俊俋俏俐俗保俞俟俠信俦俩俪俭修俱俸俾倆倉倍倏候倚倡倢倧倨倩倪倫值偃偅假偈偉偌做健偩偮偵偶偿傅傌傑備傢傧储傭傲傳僅僇僎像僑僧僵價僾儀儂億儆儉儌儐儒儔償優儲儷兄光克全兰关兹冉再冏农几凡凤凯凰凱加勗募包化区區华单卖卢卧危卷叠口古句另只召可台史右叶号司各合吉同名后向吕吟含启吳吴吾呂呈告员周味呻命和咕咖咨咸品哈哑員哥哲唐唤售唯唱商啞啟啤善喉喚喜喧喬單嗣嘉嘱器嚥嚴囑回因团园固国图國園圖團圮场坚坤埏堂堅堠堤場增士壹壽处备复夢夭央奈如妃委姬姿娅娇娟婉婚婭媛媴嬌存孙季孫宗宛客宣宸容宾富寿小少尖尚尤尼尾屯岱崇川巡巫已巷巽帆幅幕平庄庭庸延廷建开异弓引弗弘弟张弥弦弭弯弱張強强弼彌彎当彪彭彷役彻往律徐徕得從徠御復循微徵德徹思悠慕慢懂战戰扬拒招挡捐捺提揚搭撰擋攸敦日旦旨早旭时旺旻昀昂昆昇昊昌明昏易昙星映春昧昭是昱昶显晁時晃晉晋晏晒晓晔晖晚晟晤晨普景晰晴晶智暄暉暖暗暢暮暹曄曆曇曉曙曜曝曦曬書曹曼曾替會朵杏来杨杰柏柯柰桐档桥桦梁梦棕棻楊榬槽樺橋檔次欢欤欧欲欷欸款歇歌歐歟歡毕民氾汎沣河治泓泛泡涎渭游湾溒溶漠漫潘潛潜澎灃灣炅炬炯烯照熔熙猿珅琬瑛瓷生甫田由甲申男甸町畅界畏畔留畢略畦番異畴當畸疆疇疊皮皱皺益监監盧知砷示礼礽社祀祁祅祇祈祉祐祕祖祜祝神祟祠祥祭祯祷祸祺祿禁禄禅禍禎福禦禧禪禮禱禳禹禾秀私秉秋种科秒秘租秦秩积称移稀稅程稍税稚稜種稱稳稷稹稼稽稿穆積穗穠穩穫竖米类粟粮粽精糕糖糧素細紳紹結給綜繒绅细绍结给综缯署羡羨耀胡脩臣臥臨臺臻艺艾芃芊芙芝芬芮花芳芷芸芹苍苏苑苒苓苔苗苣若苯英苹苻茂范茉茗茜茱茲茵茶茹荀荃荆草荊荏荐荞荡荷荻莅莆莉莊莎莒莓莘莞莩莫莱莲菁菇菊菖菩華菱菲菽萊萌萍萝萤萧萨萬萱萲落葉著葛董葵蒂蒋蒙蒞蒨蒲蒼蓁蓄蓉蓓蓝蓟蓦蓨蓬蓮蔓蔘蔚蔡蔣蔽蕃蕉蕎蕩蕭蕴薆薇薊薑薛薦薩薪薰薹藉藍藏藝藤藩藻蘇蘊蘋蘩蘭蘴蘿虎虔處虚虛號虫虹蛇蛉蛟蛮蛻蜀蜂蜒蜕蜜蜻蝉蝴蝶螈融螢蟬蠻袁袋袍袖襄視视訊誕誘諮諷謂譁譔讯讽诞诱谓谘谷豆豎豐貝貞財貢責貴貸費貿賀資賈賓賜賢賣賦賴贊贝贞贡财责贤贵贷贸费贺贾资赋赐赖赞起超距躬輝轅辉辕辰農边达迁迅过迈迎运近还进远违连迟迦迪述迷迺送适选逊逑递途逖通速造逢連週進逵逸逻遁遊運過道達違遙遜遞遠遣遥適遲遷選邀邁還邈邊邏邑邓邝邢那邦邬邱邵邹郁郎郝郭都鄉鄒鄔鄞鄧鄺酒酥醒鈿鉀鉅錼鍚鎔鎱鏵鑑钜钾钿铧門閃開閎閑間閔閨閩闆闕關门闪闲闳间闵闺闽阙阳际陈陳陽際靡頁頂項順須頌預頗頡頤題顏顛類顧顯页顶项顺须顾颂预颇颉颐题颜颠風风饌馔香驀驊骅高魯鲁鳳鳴鵑鸣鹃麥麦麵黃黄黍黎鼓龍龙`,
	},
	ZodiacRabbit: {
		Name: ZodiacRabbit,
		Ji:   `主丽乌习书买乾人亿仁仅仇今介仑仕付仙仟代令以仪仰仲仳仵件价任份企伂伊伍休优伙会伟传伥伦伯估伴伶伸伺似但佈位低住佐佑何余佛作佟佣佩佰佳佶佼佾使侄來例侍侑侖供依侠侦侨侬侯係促俊俋俐俗保俟俠信俩俪俭修俱俸俾倀倆倍倏候倚倡倢倨倩倪倫值偃偅假偈偉偌做健偩偮偵偶偿傅傌傑備傢傧储傭傲傳僅僇像僑僧僵價僾儀儂億儆儉儌儐儒償優儲儷光全农冠凤凰勗勝化医卖卧双君命咙唬唱嘱嚁嚥嚨囑场坏堤場增壞备大天太夭头婚婴媜嬰存宇安宝实宣宪宸宾寥實寶尤就山岗岱崔崗崷嵺嶍巫帐帝帳庆应庞廖张張录彪心必忆忌忍志忙忠快念忻忽怀态怅怎怔怕怖怜思怡急性怨怪恃恆恋恐恕恢恣恩恬恭息恰恳悄悅悉悍悒悔悖悟悠悦悬悲悵悸悽情惜惟惠惢惧惭惯想惶愁愃愈愉愍意愚愛感愠愤愧愫愿慄慇慈態慌慍慎慕慚慢慣慧慨慰慶慷慾憀憎憐憤憩憲憶憾懂懇應懋懷懸懼懿戀扬振提揚摎攸日旦旨早旬旭时旺旻昀昂昆昇昊昌明昏易昙星映春昧昭是昱显晁時晃晉晋晏晒晓晖晚晟晤晨普景晰晴晶智暄暉暑暖暗暢暮曆曇曉曙曜曝曦曬書曹曼曾會有朔望朝期朧来杨杰柏栖栩桢楊楨槢槽樛欢欤欧欲欷欸款歇歐歙歟歡氏民泷浓涨漉漫漲漻潁潛潜濃瀚瀧炅烏烯照煪熠熤燿爱玉王玖玛玥玫玮环现玲玺珀珊珍珑珝珠班珮珺現球琅理琉琏琐琦琨琪琳琴琼瑋瑙瑚瑛瑜瑞瑟瑣瑤瑩瑪瑰瑶璀璃璇璉璋璎璞璧璨環璽璿瓊瓏瓔甯畅矇碧祀祯禎秘稚穎穠笼簏籠素績繆繒纓纖纤绩缨缪缯署羽羿翁翃翅翊翌翎習翔翕翘翟翠翡翦翮翰翱翳翹翻翼耀聆肯育胖胜胞胡胤胥胧胭胶能脈脉脩腾膠膳臆臥苓苹莫莹莺蓨蔍蕙薑藋蘋袋袖袭裴褶襲西要覃詡譾诩谫豂貝貞財責貴買貸費貽賀賂賄資賈賓賜賠賣質賴賺購賽贝贞财责质购贵贷费贺贻贾贿赂资赐赔赖赚赛趯跃蹟躍輝轆轇辉辘辰農遒遗遺邬郁郑郦鄔鄭酈酉酊酋配酒酗酥酪酬醇醋醍醐醑醒醫醮金釧鈞鈴鉀鉻銀銮銳鋒鋹鋼錄錢錦錯鍊鍚鍵鎮鏡鐘鐵鑑鑽鑾钏钟钢钧钱钾铁铃铬银锋锐错锦键镇镜長长阳陇陽隴雀雄集雙零震青韔順領頤頭題顏願顟顧顯顺顾领颍颐颖题颜飂飛飞駐騛騰驻魯鲁鳥鳳鳴鴒鴗鴛鴦鴻鴿鵑鵠鵡鵬鵰鵲鶖鶧鶯鶴鶼鷚鷥鷲鷸鷹鷺鸐鸕鸚鸝鸞鸟鸣鸬鸯鸳鸶鸽鸾鸿鹂鹃鹄鹉鹊鹏鹣鹤鹦鹨鹫鹬鹭鹰鹿麒麓麗麝麟齡龄龍龐龔龙龚`,
	},
	ZodiacDragon: {
		Name: ZodiacDragon,
		Ji:   `万专东严丬丰临义乔乡二亚些亞亭亮人亿什仁仅仇介仑仕付仙仟代令以仪仰仲仳仵件价任份企伂伊伍休优伙伟传伦伯估伴伶伸伺似佃但佈佌位低住佐佑佔何余佛作佟佣佩佰佳佶佼佾使侄來例侍侑侖供依侦侨侬侯係促俊俋俐俗保俟信俩俪俭修俱俸俾倆倍倏候倚倢倨倩倪倫值偃偅假偈偉偌偏做健偩偮偵偶偿傅傌傑備傢傧储傭傮傲傳僅僇像僑僧僵價僾儀儂億儆儉儌儐儒償優儲儷允元兄充先克免兔兢全兰关兹养兽冈冉冊册再冏冕农冤冻准凍减刀刊刑划列刘则刚创删判別利刪别到制刷券刻剀剂則剋前剑剛剧副剴創劃劇劉劍劑劝功加务动助励劳势勇勉動勘務勞募勢勤勵勸匕化匠匡匮匯匱匹区匿區卉华协卓協单卖博占卢卧卯印卿厂厅厢厦厨叠口古句另召可台叱史右叶号司各合吉同名后向吕吟吠含启吳吴吾呀呂呈告员周味命和咕咖咨咸品哈員哥哭哲唐唤售唯唱商問啟啤善喚喜喧喬單嗣嘈嘉嘱器嚴囑回因团园国图圃國園圖團坏坚城埔域堂堅塞壞士壮壯壹壽处备夢夭央奇奖如姍委姗姜姬姿威娅娇娟娩婭嬌孱宁它宅安宋完宏宓官宙宛宜宝实宠审客宣室宥宪宫宬宮宰宴家宸容宽宾宿寂寄寅密寇富寓寞察寥實寧審寬寰寵寶寻寿專尋小少尖尚尤就尼尾局层居屆屈届屋屏属屠屡屢屣層履屬屯山屹岌岑岗岚岛岡岩岭岱岳岷岸峋峒峙峡峦峨峰島峸峻峽崆崇崑崔崗崙崧崴嵐嵘嵩嶸嶺嶽巍巒巖川州巡巢工巧巫巳巴巷巽席幅幕干平广庄庆庇序库应底庖店庙庚府庞度座庫庭庵庶康庸廂廈廉廊廓廕廖廚廟廠廣廩廪廳延廷建开异式弓引弗弘弟张弥弦弭弯弱張強强弼彌彎当彗彧彪彭心必忆忌忍志忙忠快念忻忽怀态怎怔怕怖怜思怡急性怨怪恃恆恋恐恕恢恣恩恬息恰恳悄悅悉悍悔悖悟悠悦悬悲悶悸悽情惊惑惜惟惠惧惭惯想惶愁愃愈愉愍意愚愛感愠愤愧慄慇慈態慌慍慎慕慚慢慣慧慨慰慶慷慾憎憐憤憩憲憶憾懂懇應懋懷懸懼懿戀戈戊戌戍戎戏成戕或战戡截戫戰戲戴戶户房所扁扈扉托抑拒招指挡挽捐捨搭擋支攸敦斌斤断斯新斷旨旮昂春昭晚晟暮曲曹服朗朧木朴权杆杏杜束来杰東松林枝柔柯柳柴柵栅栏桐桓档桥桦桨桿梁梦棚棫棻楚楼榕榴槳槽樓樺橋檔欄權次欢欣欤欧欲欷欸款歇歌歐歟歡此武比毕毧汇沣沪河治泓洞洲浅浔浦涛润淢淺減渭湾準滅滬演漕漠潘潤潯潺澎澜濤瀾灃灣灭炘炯炽然照熔熙熾燃爱爿牆片版牋牌牍牒牘牙牢犀犬状犹狀狂狄狐狗狘狨狩独狮狱猎猛猜猭献猴猶猷猿獂獄獅獎獨獲獵獸獻玕玡玮玼珊珋瑊瑋瑛瓷甫甯田由甲男甸界畎畏畔留畢略異畴當畸疆疇疊益监盛監盧直相矇矛矜矞矢知矩矫短矯礼禅福禪禮禾秀私秉秋种科秒秘租秦秩积称移稀稅程稍税稚稜稢種稱稳稷稹稼稽稿穆積穗穩穫穴究穷空穿突窃窈窍窑窕窗窘窜窟窥窦窮窯窺竄竅竇竊竖竿符箴篇米类粟粮精糊糕糖糟糧細紹結給絨編織纖纤细织绍绒结给编罔罕羊美羚羡羢群羨義羬羲翩聊聞聿肇肖肩胖胞胡胤胥胧胭胶能脈脉脩脯膠膳膺臆臣臥臧臨臺臻舍舒舖艮良艰艱艺艾芃芊芙芝芬花芳芸芽苍苏苑苓苔苗苣若苯英苹苻茂范茅茆茉茙茲茵茶茹荀荆草荊荏荐荞荡荷荻莅莆莉莊莎莒莓莘莞莩莫莱莲获莸菁菇菊菟菩華菱菲萊萌萍萝萤萧萬萱萲落葉葛董葴葵蒂蒋蒙蒞蒲蒼蓁蓄蓉蓓蓝蓦蓨蓬蓮蔓蔘蔚蔡蔣蔽蕃蕉蕊蕎蕕蕙蕥蕩蕭蕴薆薇薑薦薪薰藉藍藏藝藤藩藻蘇蘊蘋蘩蘭蘴蘿虎虔處號虹蜀蜜蝉蝴蝶融螢蟬袋装裕裝襄託試誘誠諮諴謂譁識识试诚诱谓谘谷豆豊豎豐豪貸費資賓賢賣質賽贤质贷费资赛赶超越趕距躬輓輔辅農达迁迈迎运近还进远连迪迷选通造逢連進逵逸運遍道達遠遭遷選邁還邑邓邝邢那邦邪邰邱邵郁郎郑郕郭都鄉鄧鄭鄺酥醴鈿鉀鉅鉚鉞銅鋸鍼鎔鎦鏵鑑钜钺钾钿铆铜铧锯镏門開閎閔閣閨閩閱闆闈闊闕闖關闡闢门问闯闱闳闵闷闺闻闽阁阅阐阔阙陀陇陈陳隇隴雉雪靡頤類颐養馗駥騙騮驀驊驚骅骗骝高魠鳴鷢鷸鷹鸣鹬鹰麥麦麵黃黄黍黎默鼓齒齡齿龄龐`,
	},
	ZodiacSnake: {
		Name: ZodiacSnake,
		Ji:   `万丘丞丬丰久乏书争亏些亥亮人亿什仁仅仇今介仍仑仓仕付仙仟代令以仪仰仲仳仵件价任份企伂伊伍休伕众优伙会伟传伦伯估伴伶伸伺似但佈佌位低住佐佑佔何佘余佛作佟佣佩佰佳佶佼佾使侄來例侍侑侖供依侠侦侨侬侯係促俊俋俏俐俗保俟俠信俩俪俭修俱俸俾倆倉倍倏候倚倡倢倨倩倪倫值偃偅假偈偉偌做健偩偮偵偶偿傅傌傑備傢傧储催傭傲傳僅僇像僑僧僵價僾儀儂億儆儉儌儐儒償優儲儷允光入全公兰兹冈冰冲决况冷冻准凉凌凍几凯凱函刀刊刑划列刘则刚创判別利别到制刷券刻剀剂則剋前剑剛剧副剴創劃劇劉劍劑劝功加务动励劳势動勗勘務勞募勢勤勳勵勸匕化北匯卉华协卓協占卢卧卿参參又及友双叔取受变叛召叱号呀命和咳唬唱嚎场堤場增壕壮壯壹处备外多够夠夢夥夭央奖委威婚嫁子孔存孙孝孟季孩孫它宣家寅密将將少尔山岂岑岗岚岛岡岩岭岱岳岸峥峰島峻崇崔崖崗崙崢崧嵐嵘嵩嶸嶺嶽巖巫幏幕干平幼幾广庄庇廊廣式引弟弥弦弭弱弼彌录彗彪彭悠愛慕慢慮懂戈戎戏戕或戡截戲戴托扬承披拯指挣据掙提揚搋摧據支收攸改攻救敝敢散数數文斤断斯新斷日旦旨早旬旭时旺旻昀昂昆昇昊昌明昏易昕昙星映春昧昭是昱昶昺晁時晃晉晋晏晒晓晔晖晚晟晤普景晰晴晶智暄暉暖暗暢暮暴曄曆曇曉曙曜曦曬書曹曼曾會朦机杆来杨杰杶枝柏染柔树核根桑桨桿梁梦棻楊槳槽樹機橡檖檬次欢欣欤欧欲欷欸款歇歌歐歟歡此步毅比水永氾求汇汉汎汐汝江池汤汪汯汰汲汴汹決沂沅沈沉沍沐沖沙沛沣沥沧沪河沸油治沼沿況泉泊泓法泛泞泠波注泰泳泽洁洋洗洛洞津洪洲洳洵洶洸洺活洽派流浅浈浊济浒浓浙浚浤浦浩浪浮海浸涂涅消涌涛涟涣涤润涨液涵涼淀淇淋淑淙淞淡淥淦淨深淳淵混淺添淼清渊渌渙渡渤温港渱游渺渼湃湄湖湛湞湟湧湯湾湿溎源準溥溪溫溱溶滄滋滌滔滚满滥滨滬滸滾滿漂漆漓演漠漢漣漪漫漲漳潔潘潛潜潢潤潭潮澄澎澤澲激濁濃濈濕濘濟濠濡濤濫濬濯濱瀚瀝瀰瀴灃灏灝灣炅炉炖炽烯照熾燧爐爭爱父爽爾爿牆片版牋牌牒牟状狀猪猭献獎獻玕玡玼球琅琉琥瑑瑛瑯璲畅皓皮益盧直眾睁睜矛矜矞矢矣知矩矫短矯祀祿禄禠禭禹禾秀私秉秋种科秒秘租秦秩积称移稀稅程稍税稚稜種稱稷稹稻稼稽稿穆積穗穟穠穿立竖端竿笋筍筝箏篆籙米粉粟粮粲精糕糖糠糧紘級素紫絃絨綱緣繒織纂级纲织绒缘缯罕署羡羨耀聚肤能脩膚臥臧臻艮良艰艱艺艾芃芊芋芙芝芥芦芬芮花芷芸芽苏苑苒苓苔苯英苹苻苾茂范茅茉茗茲茶茹荀荃荆草荊荏荞荡荷荻莅莆莉莊莎莒莓莘莞莩莫菇菊華菱菲萌萍萝萧萬萱萲落葛董葵蒂蒋蒙蒞蒨蒲蓓蓝蓦蓨蓫蓬蔓蔘蔚蔡蔣蔽蕉蕎蕥蕩蕭蕴薰藉藍藏藝藤藩蘆蘇蘊蘋蘩蘭蘴蘿虎虑虔處虚虛虞號虢虧蝝袋袖被装裘裝襐託試該誘諍變试诤该诱豆豈豎豐豚象豢豥豪豫豬貸質质贷赶趕躬輝辉迷逐逑递遂遇遞遯遽邃邪邱郎酒酥鈍鉀銀鋐鋼錄錚鍚鎵鏵鐌鐩钝钢钾铧铮银镓閎闊闳阔阳陀陟陽隧隶隸雉雙雯頓題顱顿颅题風风香驀骸魠魯鲁鴻鷸鸿鹬麥麦黃黄黍黎黑鼓齒齡齿龄`,
	},
	ZodiacHorse: {
		Name: ZodiacHorse,
	},
	ZodiacSheep: {
		Name: ZodiacSheep,
	},
	ZodiacMonkey: {
		Name: ZodiacMonkey,
	},
	ZodiacChicken: {
		Name: ZodiacChicken,
	},
	ZodiacDog: {
		Name: ZodiacDog,
	},
	ZodiacPig: {
		Name:      ZodiacPig,
		Xi:        "",
		XiRadical: "",
		Ji:        `一乙几刀力匕三上凡也土大川己已巳干弓之仁什今天太巴引戈支斤日比片牙王爿主乏代刊加功包卉占古召叱央它市布平弘弗旦玉申皮矛矢石示氾光列刑危夷妃州帆式戎早旨旭朵此虫血衣圮托汎伸佔但伯伶别判君呀坎壮希庇廷弟彤形杉系邑礽玖些依例刻券刷到制协卓卷呻坤奇奈宗宛延弦或戕旺易昌昆昂明昀昏昕昊昇枝欣版状直知社祀祁纠初采长忻玕佌旻炘炅罕泓泡泛玫表亮侯係前剋则勇垣奐宣巷帝帅建弭彦春昭映昧是星昱架柏矜祉祈祇禹穿竿紂红纪纫紇约紆羿虹衫风玡玠柰祅紈迅巡芝芽指珊玲珍珀玳倡候修刚奚孙家差席师庭时晋晏晃晁书核栩矩砷祕祐祠祟祖神祝祚纺纱纹素索纯纽级紜纳纸纷翁蚩袁袂衽讯託起躬珅倧扆祜紘紓衿衾邪邦那迎返近范茅苓班琉珮珠乾健凰副务勘曼唱唬婉婚将崇常张强彬彩彫旋晚晤晨曹勗桿烯祥票祭絃统扎绍紼细绅组终羚翌翎聆彪蛉被袒袖袍袋责玼珩埏紵紾袗邵述迦迪能情掛採捺涎浅琅球理现剴创劳喉场堤堠媛嵐巽帧弼彭斯普晰晴晶景智曾棕牌番短结绒绝紫絮丝络给绚絳翕蛟裁裂视诊费贵须邰琁珺牋矞絪絜軫陀郎郁送迷迺莫庄莉提扬琪琳琥琴琦琨勤势园戡暗暉暖暄会杨枫照煜牒祺禄禁经绢绥群蜀蛾蜕蜂裟裙补裘装裕试钾雉零琬琰琝媴絻郡通连速造透逢途準猿瑚瑟瑞瑙瑛瑜署僎划寥廖彰截畅碧禎福祸粽绽綰综绰綾绿紧缀网纲綺绸绵綵纶维绪緇綬臧蜜蜻裳裴裹裸製褚豪宾赶闽溒瑄瑋榬綪緁緆緋綖綦蜒裱魠郭都週逸进慢漫瑶琐玛瑰剧刘剑增审影暮槽桨毅奖莹稼缔练纬緻缄缅编缘缎缓緲緹翩蝴蝶褐复褓褊诞赏质辉驻鲁齿摎漻褌褋褙陈乡运游道达违遁撰潜澎璋璃瑾璀剂勋战历晓曇炽御縑縈县縝縉萤融褪裤褫讽醒骇璇璉縕螈錼阳邹远逊遣遥递蒋璟璞励弥戏戴曙墙矫禧禪绩繆缕总纵縴縵翼襄褸辕鍚鄔蟉襁适迁环璦璨断曜璧织缮绕繚绣繒蝉顏题蕥鎱际郑邓选迟薑璿嚥曝牘璽疆祷繫绎绳绘缴襠襟识赞譔还迈邀藏琼劝曦繽继耀腾龄繾饌邈瓏樱缠续边瓔弯禳衬鷚晒缨纤襴鷸蛮纘逻湾缆驪別壯協狀糾長則帥彥紅紀紉約風剛孫師時晉書紡紗紋純紐級納紙紛訊務將張強統紮紹細紳組終責淺現創勞場幀結絨絕絲絡給絢視診費貴須莊揚勢園會楊楓祿經絹綏蛻補裝試鉀連劃暢禍綻綜綽綠緊綴網綱綢綿綸維緒賓趕閩進瑤瑣瑪劇劉劍審槳獎瑩締練緯緘緬編緣緞緩複誕賞質輝駐魯齒陳鄉運遊達違潛劑勳戰曆曉熾禦縣螢褲諷駭陽鄒遠遜遙遞蔣勵彌戲牆矯績縷總縱轅適遷環斷織繕繞繡蟬題際鄭鄧選遲禱繹繩繪繳識贊還邁瓊勸繼騰齡櫻纏續邊彎襯曬纓纖蠻邏灣纜`,
		JiRadical: "",
	},
}

var zodiacSupportList = map[string]bool{
	ZodiacRat:     true,
	ZodiacCow:     true,
	ZodiacTiger:   true,
	ZodiacRabbit:  true,
	ZodiacDragon:  true,
	ZodiacSnake:   false,
	ZodiacHorse:   false,
	ZodiacSheep:   false,
	ZodiacMonkey:  false,
	ZodiacChicken: false,
	ZodiacDog:     false,
	ZodiacPig:     true,
}

// GetZodiac ...
func GetZodiac(c chronos.Calendar) *Zodiac {
	z := chronos.GetZodiac(c.Lunar())
	b := zodiacSupportList[z]
	if !b {
		fmt.Println("[Warning]生肖:", z, ",该生肖库资料尚未补足")
	}
	if v, b := zodiacList[z]; b {
		return &v
	}
	return nil
}

func (z *Zodiac) zodiacJi(character *Character) int {
	if strings.IndexRune(z.Ji, []rune(character.Ch)[0]) != -1 {
		return -3
	}
	return 0
}

func filterZodiac(c chronos.Calendar, chars ...*Character) bool {
	return GetZodiac(c).PointCheck(3, chars...)
}

//PointCheck 检查point
func (z *Zodiac) PointCheck(limit int, chars ...*Character) bool {
	for _, c := range chars {
		if z.Point(c) < limit {
			return false
		}
	}
	return true
}

//Point 喜忌对冲，理论上喜忌都有的话，最好不要选给1，忌给0，喜给5，都没有给3
func (z *Zodiac) Point(character *Character) int {
	dp := 3
	dp += z.zodiacJi(character)
	dp += z.zodiacXi(character)
	return dp
}

func (z *Zodiac) zodiacXi(character *Character) int {
	if strings.IndexRune(z.Xi, []rune(character.Ch)[0]) != -1 {
		return 2
	}
	return 0
}
