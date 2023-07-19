package main

import "time"

type (
	Timing struct {
		Blocked         float64 `json:"blocked"`
		Dns             float64 `json:"dns"`
		Ssl             float64 `json:"ssl"`
		Connect         float64 `json:"connect"`
		Send            float64 `json:"send"`
		Wait            float64 `json:"wait"`
		Receive         float64 `json:"receive"`
		BlockedQueueing float64 `json:"_blocked_queueing"`
	}
	Entry struct {
		Initiator struct {
			Type  string `json:"type"`
			Stack struct {
				CallFrames []struct {
					FunctionName string `json:"functionName"`
					ScriptId     string `json:"scriptId"`
					Url          string `json:"url"`
					LineNumber   int    `json:"lineNumber"`
					ColumnNumber int    `json:"columnNumber"`
				} `json:"callFrames"`
				Parent struct {
					Description string `json:"description"`
					CallFrames  []struct {
						FunctionName string `json:"functionName"`
						ScriptId     string `json:"scriptId"`
						Url          string `json:"url"`
						LineNumber   int    `json:"lineNumber"`
						ColumnNumber int    `json:"columnNumber"`
					} `json:"callFrames"`
					Parent struct {
						Description string `json:"description"`
						CallFrames  []struct {
							FunctionName string `json:"functionName"`
							ScriptId     string `json:"scriptId"`
							Url          string `json:"url"`
							LineNumber   int    `json:"lineNumber"`
							ColumnNumber int    `json:"columnNumber"`
						} `json:"callFrames"`
						Parent struct {
							Description string `json:"description"`
							CallFrames  []struct {
								FunctionName string `json:"functionName"`
								ScriptId     string `json:"scriptId"`
								Url          string `json:"url"`
								LineNumber   int    `json:"lineNumber"`
								ColumnNumber int    `json:"columnNumber"`
							} `json:"callFrames"`
						} `json:"parent,omitempty"`
					} `json:"parent,omitempty"`
					ParentId struct {
						Id         string `json:"id"`
						DebuggerId string `json:"debuggerId"`
					} `json:"parentId,omitempty"`
				} `json:"parent"`
			} `json:"stack"`
		} `json:"_initiator"`
		Priority     string `json:"_priority"`
		ResourceType string `json:"_resourceType"`
		Cache        struct {
		} `json:"cache"`
		Connection string `json:"connection"`
		Pageref    string `json:"pageref"`
		Request    struct {
			Method      string `json:"method"`
			Url         string `json:"url"`
			HttpVersion string `json:"httpVersion"`
			Headers     []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"headers"`
			QueryString []interface{} `json:"queryString"`
			Cookies     []struct {
				Name     string    `json:"name"`
				Value    string    `json:"value"`
				Path     string    `json:"path"`
				Domain   string    `json:"domain"`
				Expires  time.Time `json:"expires"`
				HttpOnly bool      `json:"httpOnly"`
				Secure   bool      `json:"secure"`
				SameSite string    `json:"sameSite,omitempty"`
			} `json:"cookies"`
			HeadersSize int `json:"headersSize"`
			BodySize    int `json:"bodySize"`
			PostData    struct {
				MimeType string `json:"mimeType"`
				Text     string `json:"text"`
			} `json:"postData"`
		} `json:"request"`
		Response struct {
			Status      int    `json:"status"`
			StatusText  string `json:"statusText"`
			HttpVersion string `json:"httpVersion"`
			Headers     []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"headers"`
			Cookies []struct {
				Name     string     `json:"name"`
				Value    string     `json:"value"`
				Path     string     `json:"path"`
				Domain   string     `json:"domain"`
				Expires  *time.Time `json:"expires"`
				HttpOnly bool       `json:"httpOnly"`
				Secure   bool       `json:"secure"`
				SameSite string     `json:"sameSite"`
			} `json:"cookies"`
			Content struct {
				Size        int    `json:"size"`
				MimeType    string `json:"mimeType"`
				Compression int    `json:"compression"`
				Text        string `json:"text"`
			} `json:"content"`
			RedirectURL  string      `json:"redirectURL"`
			HeadersSize  int         `json:"headersSize"`
			BodySize     int         `json:"bodySize"`
			TransferSize int         `json:"_transferSize"`
			Error        interface{} `json:"_error"`
		} `json:"response"`
		ServerIPAddress string    `json:"serverIPAddress"`
		StartedDateTime time.Time `json:"startedDateTime"`
		Time            float64   `json:"time"`
		Timings         Timing    `json:"timings"`
	}
	HAR struct {
		Log struct {
			Version string `json:"version"`
			Creator struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"creator"`
			Pages []struct {
				StartedDateTime time.Time `json:"startedDateTime"`
				Id              string    `json:"id"`
				Title           string    `json:"title"`
				PageTimings     struct {
					OnContentLoad float64 `json:"onContentLoad"`
					OnLoad        float64 `json:"onLoad"`
				} `json:"pageTimings"`
			} `json:"pages"`
			Entries []Entry `json:"entries"`
		} `json:"log"`
	}
)

type (
	API struct {
		Url             string    `json:"url"`
		BodySize        int       `json:"bodySize"`
		StartedDateTime time.Time `json:"startedDateTime"`
		EndedDateTime   time.Time `json:"endedDateTime"`
		Time            float64   `json:"time"`
		Timings         Timing    `json:"timings"`
	}
	Result struct {
		FileName     string  `json:"fileName"`
		TotalRunTime float64 `json:"totalRunTime"`
		API          []API   `json:"api"`
	}
)
