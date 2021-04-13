// Code generated for package esp32 by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/stub_flasher.json
package esp32

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataStub_flasherJson = []byte(`{"code_start": 1074331656, "code": "0080FD3FFF0F0000364100B1FDFFA2DB4098AA7089119A98B089B09818B1F9FF973B0E1BB9B9189A8822480888BA1B8889BA1DF00800F43F0000F43F30F0FD3F081000000C00F43F1C00F43F1000F43F36410021F8FFC02000480221E9FF52D24006360000000C0681F4FFB802C02000A808A09074B65B02862D0081F0FF808BA08808A0080082A0C08719028628000C1889020C0889B5C6250082A0DB8799040C38862200C2A0C088A5C7992F1B882648024600000C0889A588A5A1E1FF7098118A892088B00C099918AA8892480088B5979803060D00009902061500661B1370A8118A8AA1D6FF2088B0AA88924800860E0000A0A07425F1FF060D0000000082A0DC87191982A0DD87191B0C43390221CDFFC020003802C020003902860A00A2A0C00601000000A2A0DBE5EDFF0C2889021B6660607437B60286CAFF61C2FFC02000380630307456A3F121C0FFC0200049021DF000000070E2FA3F1C20F43FFFFFFFDF2020F43FFFFFFF030000005C0020F43F000000400420F43F364100A1F6FF65860091F5FF81F6FFC02000A809808A10A1F4FFC020008909C02000980A81F2FF80991081F1FF808920C02000890A91F0FF81EFFFC020009908C0200098085679FFA1EDFFC02000290AC0200039082D091DF00000000000FD3F3040000000800000580009402400F43F64000094FFFF00000000800000000100000000017CDA05404CC40040EC67004008680040500006409CDA0540C8C200401CDB0540FC670040366102A2C15822613C42613F81F4FFE0080042213C3C224040B456442A3040B43C3256C429A594004D0A3C42562A29516BFFC1E1FFBD0A50A52081E9FFE00800C1DFFFA1DCFF40B42081E5FFE00800B1DCFFCD050C5A81E3FFE0080021DAFF2C0AC02000680262613D61D7FFC020006902616BFF22A101C02000722600202720C02000290681D8FFE00800A2A0B0BD042C4CAAA122A0D481D2FFE008002A211C8CBD04AD0281CEFFE0080062A0F06A611C8BAD026912A5610060EA03626130066E00000060EA03460D00000081C0FF4073C077B81BA2213CA070F4DC27B1BCFF65E7FF568A1771BBFF7A44460400000000B1B9FFA2213C25E6FF566A1642D41082213F56780370EA0382212F6068C070668062612FA0EA0362D51062260388156A8862D52068566A9862D530887668058A8992D54089127D098899860200780272D71077348506EFFFC02000B8A7871B0666460C464400006646ED86420000000060EA0372212CA077C06A6762612C60EA0362613E7068118A76A128FF5077B08BF7D817AA77720700076728718AFFB18BFF0C1ECD0F70A7208261429261438193FFE00800DD0A822142922143660A02862E00FD0770EA03A2212DB2213ECD0DB0AAC07A7ABD0FA2C15872612D826142926143D26141F261408185FFE0080070EA03A0EA03B2212FBAAA707AC072612F70EA03D22141F22140A2213CCD0DF0BF20A58000822142922143D22141563A06B0EA03C2212E8A665066B0A916707CC06899BA7772612E1B7662C6FD60A79362213C5C8CDA6662613C6802A999DAD6CAB1AD01D902816BFFE00800BD018BA2816AFFE008001C8BAD02654700C606003C528612003C624611003C720610004C02C60E00003C82460D00000062220037B602C68FFF22A0B0102280B2C158A2C214815AFFE0080020EA03322130A2A0B03022C0B2A0241AAA2261306542000C023145FF42213D2C0AC0200049038150FFE008001DF0001000006810000058100000701000004093004036412161FDFF72D1101A66590691F7FF0C0662671A5C2847B902862500AD078139FFE00800461600006083C0805463CD05BD01AD02A57D0056CA06BD05AD0181EFFFE00800CD05BD01AD078133FFE008005A225A6637360F51E6FF0C4B5AA1E53A00264A1606110051E4FF82271A1A5558058086C05738B006F7FF0082271A87362F82271A3738CCA1DBFFBD071AAA8124FFE00800A1D7FF1C0B1AAA81D8FFE008000C08060300005C388601005C484600005C582D081DF0B010000036C12161CBFF8CA452A06247B60286240040642051C9FF5A51AD05810BFFE00800461700A2D1108108FFE0080060736370C72010B12020A220A571008C6A52A063C617000000CD07BD01AD058104FFE00800AC74CD07BD01A2D1108100FFE00800A1E6FFB2D11010AA8081FEFEE00800A1E3FF1C0B1AAA81B2FFE008007A227033C056E3F9B1ACFFA1DDFFBAB11AAA81F5FEE00800A1D9FF1C0B1AAA81A9FFE008005D032D051DF00030F43F000000108030F43FFFFFFF003661000C0281FAFF290121FAFF809820C020002908208220C0200028098022105642FF81F5FF0C4BC02000980881F3FFAD0180891089018194FFE008001DF0002000F43F00000200FFFFFDFF44F0FD3F36810022A05531FAFF224115C02000880321F8FF0C1B202820C020002903C02000880321F5FFA2C115202810C020002903652100261A024641008201157CF20B882241148080740C9287B202063A0021EBFF2088A02808A00200B2A01010A120651E004C1226AA02463200C821B811A80125AFFF06070000B2A01010A120A51C005C1226BA02C62A00D831C821B811A801A5DAFFA24114462700B2A01010A120651A0022A06126AA02062200C821B811A80165E4FFC6F6FF000065EFFFC6F4FF000000A5470086F2FF00000022A000B2A001A2C114224114A5150022011590000000B2A01010A120E51500A0BA20AD01460700B2A01010A120E5140022AF91664A322801BD0AC020002802A2C110226104251200060600B2A01010A120A5120022AFA1668A0E8221012801C0200089020C02C6FFFF2241140C1BA2C114650F0022011582C2FA808074B6280206AFFF1DF0000000000009401027000050C300007C680040588600406C2A064038320640348500404C82004036610021F6FF0C0A4802381281F7FFE0080081F6FFE008000C0B81F5FFE00800F166FED123FFC166FEB166FEE2A1000C0A0C0281F0FFE008002901271315A1E8FF81EDFFE00800CD03BD0420A220E57700A26100A1E4FF81E8FFE00800B2A00410A120A5050025E2FFA02A20A1DDFF81E2FFE0080026620581E1FFE008001DF030C0FD3F0000FD3F2CF0FD3F30C0FD3F364100A1FCFFC1FAFF0C0BA0CCC0C0C221814AFEE0080011F8FF65F5FF91F6FF82A0FF890991F5FF89091DF0364100BD03AD0281FFFEE008001DF000A492004036610042A0C081FDFFE008006D0A479AF43A324D0272A0C0460C000081F8FFE008008D0A771A3D52A0DB579A1AA90152A0DC81F2FFE008008801571A0852A0DD571A040604008D068244001B442054C04793C8860000000C0522A0C081E8FFE00800279AF62D051DF000000090E2FA3F1020F43F00000008B02106403661000C18890181FAFF820801FCD8A1F9FFB1F6FDC1F8FF06080000C02000826A00C02000C26B00C02000980B5679FFC02000980AD852909D109901980107E9DA8603000C5BAD0181EDFFE00800880107E8F088010C0289031DF000F820F43FF830F43F36610081FDFFC0200038083030245643FF81FAFFC020003228003030245633FF10B12020A22065F7FF0C12A023831DF03661000C0222610021CEFD20A22025FCFF81D1FD31D2FDC020003908C0200038085673FF0C234602000010B12020A220A5F3FF8221008703F00C021DF00000008020F43F00000020000000023641004050148C250C121DF051BAFD8225048062E24066806738EC50A52065F6FF1CF60627000000A5F8FF569AFD81F6FEA1B8FD80821047A62E91EFFF90882091EDFFC02000826A0030A32082A00876880AB80A4BAAC02000B9094B9932C32042C4E022C220C60F00809401908820C02000890A40A0144042211B848090744040748D099D04A0989381DAFF4D034088C03039A00603000000A8044A98C02000A9094B444793F10C048198FD91D4FFC020009908C020009228005669FF50A52025ECFFA6140286D7FF22A0001DF00000364100A188FDE5EAFFB2A03530A320818FFFE008008223002D0A80881189031DF000000000000004364100AD0265E8FF4D022184FF8181FDC02000390221F9FFC020002908C020002228005662FF40A42025E6FF900000000820F43F000040003661002170FD20A220A5E4FFAD02BD01E5F8FF3D020C12569A04480122A200202410AD032901E5E2FF216CFD416DFDC02000426200C020004222005664FF30A32025E1FF21EDFF81EDFFC020009802AD03808920C020008902B80122A001A5F6FFA024831DF00000364100A5E1FFA02A20561A023154FDAD03A5DDFF8158FD91DFFFAD03C020009908C0200038085673FF25DCFF1DF00C121DF00000366100614AFD514AFDC0200078065057107148FDC020005906C0200068075146FD5066105146FD505620C020005907513EFD2A64781567B7030C121DF078457052E25057C057B408CD04BD03AD0206100050C52030B32020A220A5DDFF568AFD5084C07088C26D0A0606000070C72030BBA0826100E5DBFF880156BAFB1B667A5560607450B2412AA58736DF50C4C030BBA025DAFF0C130C02A023931DF0000000200001FFFFFFF7000000700000007C2420F43FEB00007000000050BB0000700000FFFF6B0000703B0000700B000070030000702C20F43F00000400364100519FFF91EFFFC02000880561EEFF90A810510FFD7114FF979A56C020008805606810C020006905C02000880561E7FF6068208108FDC020006905C0200068085106FD50661051E2FF505620C020005908720701C02000580862AF003B77605510707074505720C02000590861DAFF066800D7E802864500C020009805606910C020006905C02000980561D4FF606920C02000690577F80286240082070161EEFCECA8C02000880571EAFC707810C020007905C02000780651E8FC50771051C4FF505720C02000590661C5FFC65000C020009805811EFF808920C020008905C02000880651DDFC50881051B9FF505820C020005906820701C02000580672AF000B8870551080807450582071B1FFC020005906C02000680751B2FF50661052A0BB505620C020005907863B00000051A9FF61ADFF47F80861ACFFE7E80261ACFFC02000690561C2FC5101FFC02000880650582081C0FCC020005906C02000680851BEFC50661051BEFC505620C020005908720701C0200058087B7762AF00707074605510505720C020005908C62200C020008805606810C020006905620701DC36C02000780561ABFC606710C020006905C60B000000C02000880561E4FE606820C02000690581A4FC720701C0200058080B7762AF00707074605510505720C020005908619AFC517CFFC0200078065057207199FCC020005906C0200068075196FC5066105196FC5056206179FFC020005907516FFFC020006905A18BFC2A54781A62A00157B702C62C00E5AAFF818DFCE172FF3CF5C18DFCD2A1FFFD086D0EC6240080721147A53C916BFFC02000D909C02000790CC02000690FC02000780F5677FFA1B9FEBD0392A01076890AC02000780A4BAA790B4BBB32C34042C4C022C24046140000215CFFC02000790CC02000D26200C02000E26800C0200028085672FF4060144042211B242050744040742D055D0441A5FE6052932D032044C03035A0C60200002A54C02000580559024B222793F1460100A6140246D9FF0C062D061DF000000000E00000F43F0000FFFF0F00CC90004036410081FCFFAD028A8200881121F8FFF62A010C0282C8142A88C02000980881F6FFA0A074808910C088119094358A2930B28240BBC281F1FFE008001DF00000", "entry": 1074333904, "num_params": 2, "params_start": 1074331648, "data": "8E000940A2000940A20009400801094053010940220709404007094062070940820709408B0709409407094094070940AA070940BB070940DF070940", "data_start": 1073606704}
`)

func dataStub_flasherJsonBytes() ([]byte, error) {
	return _dataStub_flasherJson, nil
}

func dataStub_flasherJson() (*asset, error) {
	bytes, err := dataStub_flasherJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/stub_flasher.json", size: 8591, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/stub_flasher.json": dataStub_flasherJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"stub_flasher.json": &bintree{dataStub_flasherJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}