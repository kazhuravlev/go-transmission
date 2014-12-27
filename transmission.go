package transmission_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	TRANSMISSION_HEADER_AUTH_NAME = "X-Transmission-Session-Id"
)

type Request struct {
	Method    string                 `json:"method"`
	Arguments map[string]interface{} `json:"arguments"`
	Tag       int                    `json:"tag"`
}

func (req *Request) Bytes() []byte {
	data, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	return data
}

type Arguments struct {
	Torrents []Torrent `json:"torrents"`
}

type Response struct {
	Arguments Arguments `json:"arguments"`
	Result    string    `json:"result"`
	Tag       int       `json:"tag"`
}

type TorrentSetRequest struct {
	BandwidthPriority   int             `json:"bandwidthPriority"`
	DownloadLimit       int             `json:"downloadLimit"`
	DownloadLimited     bool            `json:"downloadLimited"`
	FilesWanted         []int           `json:"files-wanted"`
	FilesUnwanted       []int           `json:"files-unwanted"`
	HonorsSessionLimits bool            `json:"honorsSessionLimits"`
	Ids                 []int           `json:"ids"`
	Location            string          `json:"location"`
	PeerLimit           int             `json:"peer-limit"`
	PriorityHigh        []int           `json:"priority-high"`
	PriorityLow         []int           `json:"priority-low"`
	PriorityNormal      []int           `json:"priority-normal"`
	QueuePosition       int             `json:"queuePosition"`
	SeedIdleLimit       int             `json:"seedIdleLimit"`
	SeedIdleMode        int             `json:"seedIdleMode"`
	SeedRatioLimit      float32         `json:"seedRatioLimit"`
	SeedRatioMode       int             `json:"seedRatioMode"`
	TrackerAdd          []string        `json:"trackerAdd"`
	TrackerRemove       []int           `json:"trackerRemove"`
	TrackerReplace      [][]interface{} `json:"trackerReplace"`
	UploadLimit         int             `json:"uploadLimit"`
	UploadLimited       bool            `json:"uploadLimited"`
}

func (t TorrentSetRequest) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"bandwidthPriority":   t.BandwidthPriority,
		"downloadLimit":       t.DownloadLimit,
		"downloadLimited":     t.DownloadLimited,
		"files-wanted":        t.FilesWanted,
		"files-unwanted":      t.FilesUnwanted,
		"honorsSessionLimits": t.HonorsSessionLimits,
		"ids":             t.Ids,
		"location":        t.Location,
		"peer-limit":      t.PeerLimit,
		"priority-high":   t.PriorityHigh,
		"priority-low":    t.PriorityLow,
		"priority-normal": t.PriorityNormal,
		"queuePosition":   t.QueuePosition,
		"seedIdleLimit":   t.SeedIdleLimit,
		"seedIdleMode":    t.SeedIdleMode,
		"seedRatioLimit":  t.SeedRatioLimit,
		"seedRatioMode":   t.SeedRatioMode,
		"trackerAdd":      t.TrackerAdd,
		"trackerRemove":   t.TrackerRemove,
		"trackerReplace":  t.TrackerReplace,
		"uploadLimit":     t.UploadLimit,
		"uploadLimited":   t.UploadLimited,
	}
}

type Torrent struct {
	ActivityDate            int        `json:"activityDate"`
	AddedDate               int        `json:"addedDate"`
	BandwidthPriority       int        `json:"bandwidthPriority"`
	Comment                 string     `json:"comment"`
	CorruptEver             int        `json:"corruptEver"`
	Creator                 string     `json:"creator"`
	DateCreated             int        `json:"dateCreated"`
	DesiredAvailable        int        `json:"desiredAvailable"`
	DoneDate                int        `json:"doneDate"`
	DownloadDir             string     `json:"downloadDir"`
	DownloadedEver          int        `json:"downloadedEver"`
	DownloadLimit           int        `json:"downloadLimit"`
	DownloadLimited         bool       `json:"downloadLimited"`
	Error                   int        `json:"error"`
	ErrorString             string     `json:"errorString"`
	Eta                     int        `json:"eta"`
	EtaIdle                 int        `json:"etaIdle"`
	Files                   []File     `json:"files"`
	FileStats               []FileStat `json:"fileStats"`
	HashString              string     `json:"hashString"`
	HaveUnchecked           int        `json:"haveUnchecked"`
	HaveValid               int        `json:"haveValid"`
	HonorsSessionLimits     bool       `json:"honorsSessionLimits"`
	Id                      int        `json:"id"`
	IsFinished              bool       `json:"isFinished"`
	IsPrivate               bool       `json:"isPrivate"`
	IsStalled               bool       `json:"isStalled"`
	LeftUntilDone           int        `json:"leftUntilDone"`
	MagnetLink              int        `json:"magnetLink"`
	ManualAnnounceTime      int        `json:"manualAnnounceTime"`
	MaxConnectedPeers       int        `json:"maxConnectedPeers"`
	MetadataPercentComplete float32    `json:"metadataPercentComplete"`
	Name                    string     `json:"name"`
	PeerLimit               int        `json:"peer-limit"`
	Peers                   []Peer     `json:"peers"`
	PeersConnected          int        `json:"peersConnected"`
	PeersFrom               struct {
		FromCache    int `json:"fromCache"`
		FromDht      int `json:"fromDht"`
		FromIncoming int `json:"fromIncoming"`
		FromLpd      int `json:"fromLpd"`
		FromLtep     int `json:"fromLtep"`
		FromPex      int `json:"fromPex"`
		FromTracker  int `json:"fromTracker"`
	} `json:"peersFrom"`
	PeersGettingFromUs  int           `json:"peersGettingFromUs"`
	PeersSendingToUs    int           `json:"peersSendingToUs"`
	PercentDone         float32       `json:"percentDone"`
	PieceCount          int           `json:"pieceCount"`
	Pieces              string        `json:"pieces"`
	PieceSize           int           `json:"pieceSize"`
	Priorities          []int         `json:"priorities"`
	QueuePosition       int           `json:"queuePosition"`
	RateDownload        int           `json:"rateDownload "`
	RateUpload          int           `json:"rateUpload "`
	RecheckProgress     float32       `json:"recheckProgress"`
	SecondsDownloading  int           `json:"secondsDownloading"`
	SecondsSeeding      int           `json:"secondsSeeding"`
	SeedIdleLimit       int           `json:"seedIdleLimit"`
	SeedIdleMode        int           `json:"seedIdleMode"`
	SeedRatioLimit      float32       `json:"seedRatioLimit"`
	SeedRatioMode       int           `json:"seedRatioMode"`
	SizeWhenDone        int           `json:"sizeWhenDone"`
	StartDate           int           `json:"startDate"`
	Status              int           `json:"status"`
	TorrentFile         string        `json:"torrentFile"`
	TotalSize           int           `json:"totalSize"`
	Trackers            []Tracker     `json:"trackers"`
	TrackerStats        []TrackerStat `json:"trackerStats"`
	UploadedEver        int           `json:"uploadedEver"`
	UploadLimit         int           `json:"uploadLimit"`
	UploadLimited       bool          `json:"uploadLimited"`
	UploadRatio         float32       `json:"uploadRatio"`
	Wanted              []int         `json:"wanted"`
	Webseeds            string        `json:"webseeds"`
	WebseedsSendingToUs int           `json:"webseedsSendingToUs"`
}

type Tracker struct {
	Announce string `json:"announce"`
	Id       int    `json:"id"`
	Scrape   string `json:"scrape"`
	Tier     int    `json:"tier"`
}

type TrackerStat struct {
	Announce              string `json:"announce"`
	AnnounceState         int    `json:"announceState"`
	DownloadCount         int    `json:"downloadCount"`
	HasAnnounced          bool   `json:"hasAnnounced"`
	HasScraped            bool   `json:"hasScraped"`
	Host                  string `json:"host"`
	Id                    int    `json:"id"`
	IsBackup              bool   `json:"isBackup"`
	LastAnnouncePeerCount int    `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string `json:"lastAnnounceResult"`
	LastAnnounceStartTime int    `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      int    `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string `json:"lastScrapeResult"`
	LastScrapeStartTime   int    `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded"`
	LastScrapeTime        int    `json:"lastScrapeTime"`
	LastScrapeTimedOut    int    `json:"lastScrapeTimedOut"`
	LeecherCount          int    `json:"leecherCount"`
	NextAnnounceTime      int    `json:"nextAnnounceTime"`
	NextScrapeTime        int    `json:"nextScrapeTime"`
	Scrape                string `json:"scrape"`
	ScrapeState           int    `json:"scrapeState"`
	SeederCount           int    `json:"seederCount"`
	Tier                  int    `json:"tier"`
}

type Peer struct {
	Address            string  `json:"address"`
	ClientName         string  `json:"clientName"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	FlagStr            string  `json:"flagStr"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom "`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsIncoming         bool    `json:"isIncoming"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	IsUTP              bool    `json:"isUTP"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
	Port               int     `json:"port"`
	Progress           float32 `json:"progress"`
	RateToClient       int     `json:"rateToClient"`
	RateToPeer         int     `json:"rateToPeer"`
}

type File struct {
	BytesCompleted int    `json:"bytesCompleted"`
	Length         int    `json:"length"`
	Name           string `json:"name"`
}

type FileStat struct {
	BytesCompleted int  `json:"bytesCompleted"`
	wanted         bool `json:"wanted"`
	priority       int  `json:"priority"`
}

type TransmissionClient struct {
	Url        string
	Username   string
	Password   string
	Token      string
	TagCounter int
}

func New(url, username, password string) (*TransmissionClient, error) {
	t := &TransmissionClient{}
	t.Url = url
	t.Username = username
	t.Password = password
	t.TagCounter = 1
	err := t.UpdateToken()
	return t, err
}

func (t TransmissionClient) MakeRequest(request Request) (*http.Response, []byte, error) {
	req, err := http.NewRequest("POST", t.Url, bytes.NewBuffer(request.Bytes()))
	if err != nil {
		return nil, nil, err
	}
	req.SetBasicAuth(t.Username, t.Password)
	req.Header.Set(TRANSMISSION_HEADER_AUTH_NAME, t.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, body, nil
}

func (t *TransmissionClient) UpdateToken() error {
	resp, _, err := t.MakeRequest(Request{})
	if err != nil {
		return err
	}
	t.Token = resp.Header[TRANSMISSION_HEADER_AUTH_NAME][0]
	return nil
}

func (t *TransmissionClient) Send(req Request) (Response, error) {
	var resp_result = Response{}
	resp, body, err := t.MakeRequest(req)
	if err != nil {
		return resp_result, err
	}
	if resp.StatusCode == 409 {
		t.UpdateToken()
		return t.Send(req)
	}
	if err := json.Unmarshal(body, &resp_result); err == nil {
		return resp_result, nil
	} else {
		return resp_result, err
	}
}

func (t *TransmissionClient) CallMethod(method_name string, arguments map[string]interface{}) (Response, error) {
	var resp Response
	t.TagCounter += 1
	var req = Request{
		Method:    method_name,
		Arguments: arguments,
		Tag:       t.TagCounter,
	}
	resp, err := t.Send(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (t *TransmissionClient) GetTorrents(ids []int, fields []string) ([]Torrent, error) {
	if len(fields) == 0 {
		fields = GetTorrentJSONFieldsList()
	}
	args := map[string]interface{}{
		"fields": fields,
	}
	if len(ids) > 0 {
		args["ids"] = ids
	}
	resp, err := t.CallMethod("torrent-get", args)
	if err != nil {
		return []Torrent{}, err
	}
	return resp.Arguments.Torrents, nil
}

func (t *TransmissionClient) TorrentStart(ids []int) error {
	args := map[string]interface{}{
		"ids": ids,
	}
	_, err := t.CallMethod("torrent-start", args)
	return err
}

func (t *TransmissionClient) TorrentStartNow(ids []int) error {
	args := map[string]interface{}{
		"ids": ids,
	}
	_, err := t.CallMethod("torrent-start-now", args)
	return err
}

func (t *TransmissionClient) TorrentStop(ids []int) error {
	args := map[string]interface{}{
		"ids": ids,
	}
	_, err := t.CallMethod("torrent-stop", args)
	return err
}

func (t *TransmissionClient) TorrentVerify(ids []int) error {
	args := map[string]interface{}{
		"ids": ids,
	}
	_, err := t.CallMethod("torrent-verify", args)
	return err
}

func (t *TransmissionClient) TorrentReannounce(ids []int) error {
	args := map[string]interface{}{
		"ids": ids,
	}
	_, err := t.CallMethod("torrent-reannounce", args)
	return err
}

func (t *TransmissionClient) TorrentSet(req TorrentSetRequest) error {
	_, err := t.CallMethod("torrent-set", req.AsMap())
	return err
}

func GetTorrentJSONFieldsList() []string {
	type_obj := reflect.TypeOf(Torrent{})
	num_fields := type_obj.NumField()
	var res = []string{}
	for index := range make([]int, num_fields) {
		res = append(res, type_obj.Field(index).Tag.Get("json"))
	}
	return res
}
