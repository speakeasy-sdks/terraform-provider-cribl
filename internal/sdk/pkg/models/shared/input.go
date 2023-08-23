// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type InputType string

const (
	InputTypeInputCollection        InputType = "InputCollection"
	InputTypeInputKafka             InputType = "InputKafka"
	InputTypeInputMsk               InputType = "InputMsk"
	InputTypeInputHTTP              InputType = "InputHttp"
	InputTypeInputSplunk            InputType = "InputSplunk"
	InputTypeInputSplunkSearch      InputType = "InputSplunkSearch"
	InputTypeInputSplunkHec         InputType = "InputSplunkHec"
	InputTypeInputAzureBlob         InputType = "InputAzureBlob"
	InputTypeInputElastic           InputType = "InputElastic"
	InputTypeInputConfluentCloud    InputType = "InputConfluentCloud"
	InputTypeInputGrafana           InputType = "InputGrafana"
	InputTypeInputLoki              InputType = "InputLoki"
	InputTypeInputPrometheusRw      InputType = "InputPrometheusRw"
	InputTypeInputPrometheus        InputType = "InputPrometheus"
	InputTypeInputEdgePrometheus    InputType = "InputEdgePrometheus"
	InputTypeInputOffice365Mgmt     InputType = "InputOffice365Mgmt"
	InputTypeInputOffice365Service  InputType = "InputOffice365Service"
	InputTypeInputOffice365MsgTrace InputType = "InputOffice365MsgTrace"
	InputTypeInputEventhub          InputType = "InputEventhub"
	InputTypeInputExec              InputType = "InputExec"
	InputTypeInputFirehose          InputType = "InputFirehose"
	InputTypeInputGooglePubsub      InputType = "InputGooglePubsub"
	InputTypeInputCribl             InputType = "InputCribl"
	InputTypeInputCriblTCP          InputType = "InputCriblTcp"
	InputTypeInputCriblHTTP         InputType = "InputCriblHttp"
	InputTypeInputTcpjson           InputType = "InputTcpjson"
	InputTypeInputSystemMetrics     InputType = "InputSystemMetrics"
	InputTypeInputSystemState       InputType = "InputSystemState"
	InputTypeInputKubeMetrics       InputType = "InputKubeMetrics"
	InputTypeInputKubeLogs          InputType = "InputKubeLogs"
	InputTypeInputKubeEvents        InputType = "InputKubeEvents"
	InputTypeInputWindowsMetrics    InputType = "InputWindowsMetrics"
	InputTypeInputCrowdstrike       InputType = "InputCrowdstrike"
	InputTypeInputDatadogAgent      InputType = "InputDatadogAgent"
	InputTypeInputDatagen           InputType = "InputDatagen"
	InputTypeInputHTTPRaw           InputType = "InputHttpRaw"
	InputTypeInputKinesis           InputType = "InputKinesis"
	InputTypeInputCriblmetrics      InputType = "InputCriblmetrics"
	InputTypeInputMetrics           InputType = "InputMetrics"
	InputTypeInputS3                InputType = "InputS3"
	InputTypeInputSnmp              InputType = "InputSnmp"
	InputTypeInputOpenTelemetry     InputType = "InputOpenTelemetry"
	InputTypeInputSqs               InputType = "InputSqs"
	InputTypeInputSyslog            InputType = "InputSyslog"
	InputTypeInputFile              InputType = "InputFile"
	InputTypeInputTCP               InputType = "InputTcp"
	InputTypeInputAppscope          InputType = "InputAppscope"
	InputTypeInputWef               InputType = "InputWef"
	InputTypeInputWinEventLogs      InputType = "InputWinEventLogs"
	InputTypeInputRawUDP            InputType = "InputRawUdp"
	InputTypeInputJournalFiles      InputType = "InputJournalFiles"
)

type Input struct {
	InputCollection        *InputCollection
	InputKafka             *InputKafka
	InputMsk               *InputMsk
	InputHTTP              *InputHTTP
	InputSplunk            *InputSplunk
	InputSplunkSearch      *InputSplunkSearch
	InputSplunkHec         *InputSplunkHec
	InputAzureBlob         *InputAzureBlob
	InputElastic           *InputElastic
	InputConfluentCloud    *InputConfluentCloud
	InputGrafana           *InputGrafana
	InputLoki              *InputLoki
	InputPrometheusRw      *InputPrometheusRw
	InputPrometheus        *InputPrometheus
	InputEdgePrometheus    *InputEdgePrometheus
	InputOffice365Mgmt     *InputOffice365Mgmt
	InputOffice365Service  *InputOffice365Service
	InputOffice365MsgTrace *InputOffice365MsgTrace
	InputEventhub          *InputEventhub
	InputExec              *InputExec
	InputFirehose          *InputFirehose
	InputGooglePubsub      *InputGooglePubsub
	InputCribl             *InputCribl
	InputCriblTCP          *InputCriblTCP
	InputCriblHTTP         *InputCriblHTTP
	InputTcpjson           *InputTcpjson
	InputSystemMetrics     *InputSystemMetrics
	InputSystemState       *InputSystemState
	InputKubeMetrics       *InputKubeMetrics
	InputKubeLogs          *InputKubeLogs
	InputKubeEvents        *InputKubeEvents
	InputWindowsMetrics    *InputWindowsMetrics
	InputCrowdstrike       *InputCrowdstrike
	InputDatadogAgent      *InputDatadogAgent
	InputDatagen           *InputDatagen
	InputHTTPRaw           *InputHTTPRaw
	InputKinesis           *InputKinesis
	InputCriblmetrics      *InputCriblmetrics
	InputMetrics           *InputMetrics
	InputS3                *InputS3
	InputSnmp              *InputSnmp
	InputOpenTelemetry     *InputOpenTelemetry
	InputSqs               *InputSqs
	InputSyslog            *InputSyslog
	InputFile              *InputFile
	InputTCP               *InputTCP
	InputAppscope          *InputAppscope
	InputWef               *InputWef
	InputWinEventLogs      *InputWinEventLogs
	InputRawUDP            *InputRawUDP
	InputJournalFiles      *InputJournalFiles

	Type InputType
}

func CreateInputInputCollection(inputCollection InputCollection) Input {
	typ := InputTypeInputCollection

	return Input{
		InputCollection: &inputCollection,
		Type:            typ,
	}
}

func CreateInputInputKafka(inputKafka InputKafka) Input {
	typ := InputTypeInputKafka

	return Input{
		InputKafka: &inputKafka,
		Type:       typ,
	}
}

func CreateInputInputMsk(inputMsk InputMsk) Input {
	typ := InputTypeInputMsk

	return Input{
		InputMsk: &inputMsk,
		Type:     typ,
	}
}

func CreateInputInputHTTP(inputHTTP InputHTTP) Input {
	typ := InputTypeInputHTTP

	return Input{
		InputHTTP: &inputHTTP,
		Type:      typ,
	}
}

func CreateInputInputSplunk(inputSplunk InputSplunk) Input {
	typ := InputTypeInputSplunk

	return Input{
		InputSplunk: &inputSplunk,
		Type:        typ,
	}
}

func CreateInputInputSplunkSearch(inputSplunkSearch InputSplunkSearch) Input {
	typ := InputTypeInputSplunkSearch

	return Input{
		InputSplunkSearch: &inputSplunkSearch,
		Type:              typ,
	}
}

func CreateInputInputSplunkHec(inputSplunkHec InputSplunkHec) Input {
	typ := InputTypeInputSplunkHec

	return Input{
		InputSplunkHec: &inputSplunkHec,
		Type:           typ,
	}
}

func CreateInputInputAzureBlob(inputAzureBlob InputAzureBlob) Input {
	typ := InputTypeInputAzureBlob

	return Input{
		InputAzureBlob: &inputAzureBlob,
		Type:           typ,
	}
}

func CreateInputInputElastic(inputElastic InputElastic) Input {
	typ := InputTypeInputElastic

	return Input{
		InputElastic: &inputElastic,
		Type:         typ,
	}
}

func CreateInputInputConfluentCloud(inputConfluentCloud InputConfluentCloud) Input {
	typ := InputTypeInputConfluentCloud

	return Input{
		InputConfluentCloud: &inputConfluentCloud,
		Type:                typ,
	}
}

func CreateInputInputGrafana(inputGrafana InputGrafana) Input {
	typ := InputTypeInputGrafana

	return Input{
		InputGrafana: &inputGrafana,
		Type:         typ,
	}
}

func CreateInputInputLoki(inputLoki InputLoki) Input {
	typ := InputTypeInputLoki

	return Input{
		InputLoki: &inputLoki,
		Type:      typ,
	}
}

func CreateInputInputPrometheusRw(inputPrometheusRw InputPrometheusRw) Input {
	typ := InputTypeInputPrometheusRw

	return Input{
		InputPrometheusRw: &inputPrometheusRw,
		Type:              typ,
	}
}

func CreateInputInputPrometheus(inputPrometheus InputPrometheus) Input {
	typ := InputTypeInputPrometheus

	return Input{
		InputPrometheus: &inputPrometheus,
		Type:            typ,
	}
}

func CreateInputInputEdgePrometheus(inputEdgePrometheus InputEdgePrometheus) Input {
	typ := InputTypeInputEdgePrometheus

	return Input{
		InputEdgePrometheus: &inputEdgePrometheus,
		Type:                typ,
	}
}

func CreateInputInputOffice365Mgmt(inputOffice365Mgmt InputOffice365Mgmt) Input {
	typ := InputTypeInputOffice365Mgmt

	return Input{
		InputOffice365Mgmt: &inputOffice365Mgmt,
		Type:               typ,
	}
}

func CreateInputInputOffice365Service(inputOffice365Service InputOffice365Service) Input {
	typ := InputTypeInputOffice365Service

	return Input{
		InputOffice365Service: &inputOffice365Service,
		Type:                  typ,
	}
}

func CreateInputInputOffice365MsgTrace(inputOffice365MsgTrace InputOffice365MsgTrace) Input {
	typ := InputTypeInputOffice365MsgTrace

	return Input{
		InputOffice365MsgTrace: &inputOffice365MsgTrace,
		Type:                   typ,
	}
}

func CreateInputInputEventhub(inputEventhub InputEventhub) Input {
	typ := InputTypeInputEventhub

	return Input{
		InputEventhub: &inputEventhub,
		Type:          typ,
	}
}

func CreateInputInputExec(inputExec InputExec) Input {
	typ := InputTypeInputExec

	return Input{
		InputExec: &inputExec,
		Type:      typ,
	}
}

func CreateInputInputFirehose(inputFirehose InputFirehose) Input {
	typ := InputTypeInputFirehose

	return Input{
		InputFirehose: &inputFirehose,
		Type:          typ,
	}
}

func CreateInputInputGooglePubsub(inputGooglePubsub InputGooglePubsub) Input {
	typ := InputTypeInputGooglePubsub

	return Input{
		InputGooglePubsub: &inputGooglePubsub,
		Type:              typ,
	}
}

func CreateInputInputCribl(inputCribl InputCribl) Input {
	typ := InputTypeInputCribl

	return Input{
		InputCribl: &inputCribl,
		Type:       typ,
	}
}

func CreateInputInputCriblTCP(inputCriblTCP InputCriblTCP) Input {
	typ := InputTypeInputCriblTCP

	return Input{
		InputCriblTCP: &inputCriblTCP,
		Type:          typ,
	}
}

func CreateInputInputCriblHTTP(inputCriblHTTP InputCriblHTTP) Input {
	typ := InputTypeInputCriblHTTP

	return Input{
		InputCriblHTTP: &inputCriblHTTP,
		Type:           typ,
	}
}

func CreateInputInputTcpjson(inputTcpjson InputTcpjson) Input {
	typ := InputTypeInputTcpjson

	return Input{
		InputTcpjson: &inputTcpjson,
		Type:         typ,
	}
}

func CreateInputInputSystemMetrics(inputSystemMetrics InputSystemMetrics) Input {
	typ := InputTypeInputSystemMetrics

	return Input{
		InputSystemMetrics: &inputSystemMetrics,
		Type:               typ,
	}
}

func CreateInputInputSystemState(inputSystemState InputSystemState) Input {
	typ := InputTypeInputSystemState

	return Input{
		InputSystemState: &inputSystemState,
		Type:             typ,
	}
}

func CreateInputInputKubeMetrics(inputKubeMetrics InputKubeMetrics) Input {
	typ := InputTypeInputKubeMetrics

	return Input{
		InputKubeMetrics: &inputKubeMetrics,
		Type:             typ,
	}
}

func CreateInputInputKubeLogs(inputKubeLogs InputKubeLogs) Input {
	typ := InputTypeInputKubeLogs

	return Input{
		InputKubeLogs: &inputKubeLogs,
		Type:          typ,
	}
}

func CreateInputInputKubeEvents(inputKubeEvents InputKubeEvents) Input {
	typ := InputTypeInputKubeEvents

	return Input{
		InputKubeEvents: &inputKubeEvents,
		Type:            typ,
	}
}

func CreateInputInputWindowsMetrics(inputWindowsMetrics InputWindowsMetrics) Input {
	typ := InputTypeInputWindowsMetrics

	return Input{
		InputWindowsMetrics: &inputWindowsMetrics,
		Type:                typ,
	}
}

func CreateInputInputCrowdstrike(inputCrowdstrike InputCrowdstrike) Input {
	typ := InputTypeInputCrowdstrike

	return Input{
		InputCrowdstrike: &inputCrowdstrike,
		Type:             typ,
	}
}

func CreateInputInputDatadogAgent(inputDatadogAgent InputDatadogAgent) Input {
	typ := InputTypeInputDatadogAgent

	return Input{
		InputDatadogAgent: &inputDatadogAgent,
		Type:              typ,
	}
}

func CreateInputInputDatagen(inputDatagen InputDatagen) Input {
	typ := InputTypeInputDatagen

	return Input{
		InputDatagen: &inputDatagen,
		Type:         typ,
	}
}

func CreateInputInputHTTPRaw(inputHTTPRaw InputHTTPRaw) Input {
	typ := InputTypeInputHTTPRaw

	return Input{
		InputHTTPRaw: &inputHTTPRaw,
		Type:         typ,
	}
}

func CreateInputInputKinesis(inputKinesis InputKinesis) Input {
	typ := InputTypeInputKinesis

	return Input{
		InputKinesis: &inputKinesis,
		Type:         typ,
	}
}

func CreateInputInputCriblmetrics(inputCriblmetrics InputCriblmetrics) Input {
	typ := InputTypeInputCriblmetrics

	return Input{
		InputCriblmetrics: &inputCriblmetrics,
		Type:              typ,
	}
}

func CreateInputInputMetrics(inputMetrics InputMetrics) Input {
	typ := InputTypeInputMetrics

	return Input{
		InputMetrics: &inputMetrics,
		Type:         typ,
	}
}

func CreateInputInputS3(inputS3 InputS3) Input {
	typ := InputTypeInputS3

	return Input{
		InputS3: &inputS3,
		Type:    typ,
	}
}

func CreateInputInputSnmp(inputSnmp InputSnmp) Input {
	typ := InputTypeInputSnmp

	return Input{
		InputSnmp: &inputSnmp,
		Type:      typ,
	}
}

func CreateInputInputOpenTelemetry(inputOpenTelemetry InputOpenTelemetry) Input {
	typ := InputTypeInputOpenTelemetry

	return Input{
		InputOpenTelemetry: &inputOpenTelemetry,
		Type:               typ,
	}
}

func CreateInputInputSqs(inputSqs InputSqs) Input {
	typ := InputTypeInputSqs

	return Input{
		InputSqs: &inputSqs,
		Type:     typ,
	}
}

func CreateInputInputSyslog(inputSyslog InputSyslog) Input {
	typ := InputTypeInputSyslog

	return Input{
		InputSyslog: &inputSyslog,
		Type:        typ,
	}
}

func CreateInputInputFile(inputFile InputFile) Input {
	typ := InputTypeInputFile

	return Input{
		InputFile: &inputFile,
		Type:      typ,
	}
}

func CreateInputInputTCP(inputTCP InputTCP) Input {
	typ := InputTypeInputTCP

	return Input{
		InputTCP: &inputTCP,
		Type:     typ,
	}
}

func CreateInputInputAppscope(inputAppscope InputAppscope) Input {
	typ := InputTypeInputAppscope

	return Input{
		InputAppscope: &inputAppscope,
		Type:          typ,
	}
}

func CreateInputInputWef(inputWef InputWef) Input {
	typ := InputTypeInputWef

	return Input{
		InputWef: &inputWef,
		Type:     typ,
	}
}

func CreateInputInputWinEventLogs(inputWinEventLogs InputWinEventLogs) Input {
	typ := InputTypeInputWinEventLogs

	return Input{
		InputWinEventLogs: &inputWinEventLogs,
		Type:              typ,
	}
}

func CreateInputInputRawUDP(inputRawUDP InputRawUDP) Input {
	typ := InputTypeInputRawUDP

	return Input{
		InputRawUDP: &inputRawUDP,
		Type:        typ,
	}
}

func CreateInputInputJournalFiles(inputJournalFiles InputJournalFiles) Input {
	typ := InputTypeInputJournalFiles

	return Input{
		InputJournalFiles: &inputJournalFiles,
		Type:              typ,
	}
}

func (u *Input) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	inputCollection := new(InputCollection)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCollection); err == nil {
		u.InputCollection = inputCollection
		u.Type = InputTypeInputCollection
		return nil
	}

	inputKafka := new(InputKafka)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputKafka); err == nil {
		u.InputKafka = inputKafka
		u.Type = InputTypeInputKafka
		return nil
	}

	inputMsk := new(InputMsk)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputMsk); err == nil {
		u.InputMsk = inputMsk
		u.Type = InputTypeInputMsk
		return nil
	}

	inputHTTP := new(InputHTTP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputHTTP); err == nil {
		u.InputHTTP = inputHTTP
		u.Type = InputTypeInputHTTP
		return nil
	}

	inputSplunk := new(InputSplunk)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSplunk); err == nil {
		u.InputSplunk = inputSplunk
		u.Type = InputTypeInputSplunk
		return nil
	}

	inputSplunkSearch := new(InputSplunkSearch)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSplunkSearch); err == nil {
		u.InputSplunkSearch = inputSplunkSearch
		u.Type = InputTypeInputSplunkSearch
		return nil
	}

	inputSplunkHec := new(InputSplunkHec)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSplunkHec); err == nil {
		u.InputSplunkHec = inputSplunkHec
		u.Type = InputTypeInputSplunkHec
		return nil
	}

	inputAzureBlob := new(InputAzureBlob)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputAzureBlob); err == nil {
		u.InputAzureBlob = inputAzureBlob
		u.Type = InputTypeInputAzureBlob
		return nil
	}

	inputElastic := new(InputElastic)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputElastic); err == nil {
		u.InputElastic = inputElastic
		u.Type = InputTypeInputElastic
		return nil
	}

	inputConfluentCloud := new(InputConfluentCloud)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputConfluentCloud); err == nil {
		u.InputConfluentCloud = inputConfluentCloud
		u.Type = InputTypeInputConfluentCloud
		return nil
	}

	inputGrafana := new(InputGrafana)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputGrafana); err == nil {
		u.InputGrafana = inputGrafana
		u.Type = InputTypeInputGrafana
		return nil
	}

	inputLoki := new(InputLoki)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputLoki); err == nil {
		u.InputLoki = inputLoki
		u.Type = InputTypeInputLoki
		return nil
	}

	inputPrometheusRw := new(InputPrometheusRw)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputPrometheusRw); err == nil {
		u.InputPrometheusRw = inputPrometheusRw
		u.Type = InputTypeInputPrometheusRw
		return nil
	}

	inputPrometheus := new(InputPrometheus)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputPrometheus); err == nil {
		u.InputPrometheus = inputPrometheus
		u.Type = InputTypeInputPrometheus
		return nil
	}

	inputEdgePrometheus := new(InputEdgePrometheus)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputEdgePrometheus); err == nil {
		u.InputEdgePrometheus = inputEdgePrometheus
		u.Type = InputTypeInputEdgePrometheus
		return nil
	}

	inputOffice365Mgmt := new(InputOffice365Mgmt)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputOffice365Mgmt); err == nil {
		u.InputOffice365Mgmt = inputOffice365Mgmt
		u.Type = InputTypeInputOffice365Mgmt
		return nil
	}

	inputOffice365Service := new(InputOffice365Service)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputOffice365Service); err == nil {
		u.InputOffice365Service = inputOffice365Service
		u.Type = InputTypeInputOffice365Service
		return nil
	}

	inputOffice365MsgTrace := new(InputOffice365MsgTrace)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputOffice365MsgTrace); err == nil {
		u.InputOffice365MsgTrace = inputOffice365MsgTrace
		u.Type = InputTypeInputOffice365MsgTrace
		return nil
	}

	inputEventhub := new(InputEventhub)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputEventhub); err == nil {
		u.InputEventhub = inputEventhub
		u.Type = InputTypeInputEventhub
		return nil
	}

	inputExec := new(InputExec)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputExec); err == nil {
		u.InputExec = inputExec
		u.Type = InputTypeInputExec
		return nil
	}

	inputFirehose := new(InputFirehose)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputFirehose); err == nil {
		u.InputFirehose = inputFirehose
		u.Type = InputTypeInputFirehose
		return nil
	}

	inputGooglePubsub := new(InputGooglePubsub)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputGooglePubsub); err == nil {
		u.InputGooglePubsub = inputGooglePubsub
		u.Type = InputTypeInputGooglePubsub
		return nil
	}

	inputCribl := new(InputCribl)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCribl); err == nil {
		u.InputCribl = inputCribl
		u.Type = InputTypeInputCribl
		return nil
	}

	inputCriblTCP := new(InputCriblTCP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCriblTCP); err == nil {
		u.InputCriblTCP = inputCriblTCP
		u.Type = InputTypeInputCriblTCP
		return nil
	}

	inputCriblHTTP := new(InputCriblHTTP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCriblHTTP); err == nil {
		u.InputCriblHTTP = inputCriblHTTP
		u.Type = InputTypeInputCriblHTTP
		return nil
	}

	inputTcpjson := new(InputTcpjson)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputTcpjson); err == nil {
		u.InputTcpjson = inputTcpjson
		u.Type = InputTypeInputTcpjson
		return nil
	}

	inputSystemMetrics := new(InputSystemMetrics)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSystemMetrics); err == nil {
		u.InputSystemMetrics = inputSystemMetrics
		u.Type = InputTypeInputSystemMetrics
		return nil
	}

	inputSystemState := new(InputSystemState)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSystemState); err == nil {
		u.InputSystemState = inputSystemState
		u.Type = InputTypeInputSystemState
		return nil
	}

	inputKubeMetrics := new(InputKubeMetrics)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputKubeMetrics); err == nil {
		u.InputKubeMetrics = inputKubeMetrics
		u.Type = InputTypeInputKubeMetrics
		return nil
	}

	inputKubeLogs := new(InputKubeLogs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputKubeLogs); err == nil {
		u.InputKubeLogs = inputKubeLogs
		u.Type = InputTypeInputKubeLogs
		return nil
	}

	inputKubeEvents := new(InputKubeEvents)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputKubeEvents); err == nil {
		u.InputKubeEvents = inputKubeEvents
		u.Type = InputTypeInputKubeEvents
		return nil
	}

	inputWindowsMetrics := new(InputWindowsMetrics)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputWindowsMetrics); err == nil {
		u.InputWindowsMetrics = inputWindowsMetrics
		u.Type = InputTypeInputWindowsMetrics
		return nil
	}

	inputCrowdstrike := new(InputCrowdstrike)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCrowdstrike); err == nil {
		u.InputCrowdstrike = inputCrowdstrike
		u.Type = InputTypeInputCrowdstrike
		return nil
	}

	inputDatadogAgent := new(InputDatadogAgent)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputDatadogAgent); err == nil {
		u.InputDatadogAgent = inputDatadogAgent
		u.Type = InputTypeInputDatadogAgent
		return nil
	}

	inputDatagen := new(InputDatagen)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputDatagen); err == nil {
		u.InputDatagen = inputDatagen
		u.Type = InputTypeInputDatagen
		return nil
	}

	inputHTTPRaw := new(InputHTTPRaw)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputHTTPRaw); err == nil {
		u.InputHTTPRaw = inputHTTPRaw
		u.Type = InputTypeInputHTTPRaw
		return nil
	}

	inputKinesis := new(InputKinesis)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputKinesis); err == nil {
		u.InputKinesis = inputKinesis
		u.Type = InputTypeInputKinesis
		return nil
	}

	inputCriblmetrics := new(InputCriblmetrics)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputCriblmetrics); err == nil {
		u.InputCriblmetrics = inputCriblmetrics
		u.Type = InputTypeInputCriblmetrics
		return nil
	}

	inputMetrics := new(InputMetrics)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputMetrics); err == nil {
		u.InputMetrics = inputMetrics
		u.Type = InputTypeInputMetrics
		return nil
	}

	inputS3 := new(InputS3)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputS3); err == nil {
		u.InputS3 = inputS3
		u.Type = InputTypeInputS3
		return nil
	}

	inputSnmp := new(InputSnmp)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSnmp); err == nil {
		u.InputSnmp = inputSnmp
		u.Type = InputTypeInputSnmp
		return nil
	}

	inputOpenTelemetry := new(InputOpenTelemetry)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputOpenTelemetry); err == nil {
		u.InputOpenTelemetry = inputOpenTelemetry
		u.Type = InputTypeInputOpenTelemetry
		return nil
	}

	inputSqs := new(InputSqs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSqs); err == nil {
		u.InputSqs = inputSqs
		u.Type = InputTypeInputSqs
		return nil
	}

	inputSyslog := new(InputSyslog)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputSyslog); err == nil {
		u.InputSyslog = inputSyslog
		u.Type = InputTypeInputSyslog
		return nil
	}

	inputFile := new(InputFile)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputFile); err == nil {
		u.InputFile = inputFile
		u.Type = InputTypeInputFile
		return nil
	}

	inputTCP := new(InputTCP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputTCP); err == nil {
		u.InputTCP = inputTCP
		u.Type = InputTypeInputTCP
		return nil
	}

	inputAppscope := new(InputAppscope)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputAppscope); err == nil {
		u.InputAppscope = inputAppscope
		u.Type = InputTypeInputAppscope
		return nil
	}

	inputWef := new(InputWef)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputWef); err == nil {
		u.InputWef = inputWef
		u.Type = InputTypeInputWef
		return nil
	}

	inputWinEventLogs := new(InputWinEventLogs)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputWinEventLogs); err == nil {
		u.InputWinEventLogs = inputWinEventLogs
		u.Type = InputTypeInputWinEventLogs
		return nil
	}

	inputRawUDP := new(InputRawUDP)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputRawUDP); err == nil {
		u.InputRawUDP = inputRawUDP
		u.Type = InputTypeInputRawUDP
		return nil
	}

	inputJournalFiles := new(InputJournalFiles)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&inputJournalFiles); err == nil {
		u.InputJournalFiles = inputJournalFiles
		u.Type = InputTypeInputJournalFiles
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Input) MarshalJSON() ([]byte, error) {
	if u.InputCollection != nil {
		return json.Marshal(u.InputCollection)
	}

	if u.InputKafka != nil {
		return json.Marshal(u.InputKafka)
	}

	if u.InputMsk != nil {
		return json.Marshal(u.InputMsk)
	}

	if u.InputHTTP != nil {
		return json.Marshal(u.InputHTTP)
	}

	if u.InputSplunk != nil {
		return json.Marshal(u.InputSplunk)
	}

	if u.InputSplunkSearch != nil {
		return json.Marshal(u.InputSplunkSearch)
	}

	if u.InputSplunkHec != nil {
		return json.Marshal(u.InputSplunkHec)
	}

	if u.InputAzureBlob != nil {
		return json.Marshal(u.InputAzureBlob)
	}

	if u.InputElastic != nil {
		return json.Marshal(u.InputElastic)
	}

	if u.InputConfluentCloud != nil {
		return json.Marshal(u.InputConfluentCloud)
	}

	if u.InputGrafana != nil {
		return json.Marshal(u.InputGrafana)
	}

	if u.InputLoki != nil {
		return json.Marshal(u.InputLoki)
	}

	if u.InputPrometheusRw != nil {
		return json.Marshal(u.InputPrometheusRw)
	}

	if u.InputPrometheus != nil {
		return json.Marshal(u.InputPrometheus)
	}

	if u.InputEdgePrometheus != nil {
		return json.Marshal(u.InputEdgePrometheus)
	}

	if u.InputOffice365Mgmt != nil {
		return json.Marshal(u.InputOffice365Mgmt)
	}

	if u.InputOffice365Service != nil {
		return json.Marshal(u.InputOffice365Service)
	}

	if u.InputOffice365MsgTrace != nil {
		return json.Marshal(u.InputOffice365MsgTrace)
	}

	if u.InputEventhub != nil {
		return json.Marshal(u.InputEventhub)
	}

	if u.InputExec != nil {
		return json.Marshal(u.InputExec)
	}

	if u.InputFirehose != nil {
		return json.Marshal(u.InputFirehose)
	}

	if u.InputGooglePubsub != nil {
		return json.Marshal(u.InputGooglePubsub)
	}

	if u.InputCribl != nil {
		return json.Marshal(u.InputCribl)
	}

	if u.InputCriblTCP != nil {
		return json.Marshal(u.InputCriblTCP)
	}

	if u.InputCriblHTTP != nil {
		return json.Marshal(u.InputCriblHTTP)
	}

	if u.InputTcpjson != nil {
		return json.Marshal(u.InputTcpjson)
	}

	if u.InputSystemMetrics != nil {
		return json.Marshal(u.InputSystemMetrics)
	}

	if u.InputSystemState != nil {
		return json.Marshal(u.InputSystemState)
	}

	if u.InputKubeMetrics != nil {
		return json.Marshal(u.InputKubeMetrics)
	}

	if u.InputKubeLogs != nil {
		return json.Marshal(u.InputKubeLogs)
	}

	if u.InputKubeEvents != nil {
		return json.Marshal(u.InputKubeEvents)
	}

	if u.InputWindowsMetrics != nil {
		return json.Marshal(u.InputWindowsMetrics)
	}

	if u.InputCrowdstrike != nil {
		return json.Marshal(u.InputCrowdstrike)
	}

	if u.InputDatadogAgent != nil {
		return json.Marshal(u.InputDatadogAgent)
	}

	if u.InputDatagen != nil {
		return json.Marshal(u.InputDatagen)
	}

	if u.InputHTTPRaw != nil {
		return json.Marshal(u.InputHTTPRaw)
	}

	if u.InputKinesis != nil {
		return json.Marshal(u.InputKinesis)
	}

	if u.InputCriblmetrics != nil {
		return json.Marshal(u.InputCriblmetrics)
	}

	if u.InputMetrics != nil {
		return json.Marshal(u.InputMetrics)
	}

	if u.InputS3 != nil {
		return json.Marshal(u.InputS3)
	}

	if u.InputSnmp != nil {
		return json.Marshal(u.InputSnmp)
	}

	if u.InputOpenTelemetry != nil {
		return json.Marshal(u.InputOpenTelemetry)
	}

	if u.InputSqs != nil {
		return json.Marshal(u.InputSqs)
	}

	if u.InputSyslog != nil {
		return json.Marshal(u.InputSyslog)
	}

	if u.InputFile != nil {
		return json.Marshal(u.InputFile)
	}

	if u.InputTCP != nil {
		return json.Marshal(u.InputTCP)
	}

	if u.InputAppscope != nil {
		return json.Marshal(u.InputAppscope)
	}

	if u.InputWef != nil {
		return json.Marshal(u.InputWef)
	}

	if u.InputWinEventLogs != nil {
		return json.Marshal(u.InputWinEventLogs)
	}

	if u.InputRawUDP != nil {
		return json.Marshal(u.InputRawUDP)
	}

	if u.InputJournalFiles != nil {
		return json.Marshal(u.InputJournalFiles)
	}

	return nil, nil
}