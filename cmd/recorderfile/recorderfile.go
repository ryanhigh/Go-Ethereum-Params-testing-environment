package recorderfile

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var TransactionPoolInputThroughputF *csv.Writer
var TransactionPoolInputThroughputB *bufio.Writer
var NetP2pTransmissionLatencyF *csv.Writer
var PeerMessageThroughputF *csv.Writer
var DbStateWriteRateF *csv.Writer
var DbStateReadRateF *csv.Writer
var TxQueueDelayF *csv.Writer
var BlockCommitDurationStartF *csv.Writer
var BlockCommitDurationEndF *csv.Writer
var BlockValidationEfficiencyStartF *csv.Writer
var BlockValidationEfficiencyEndF *csv.Writer
var TxDelayStartF *csv.Writer
var TxDelayEndF *csv.Writer
var ContractTimeF *csv.Writer
var TxinBlockTpsF *csv.Writer
var ConsensusCliqueCostF *csv.Writer

var Workdir string
var BlockCommitDurationF *csv.Writer
var BlockValidationEfficiencyF *csv.Writer
var BlockTxConflictRateF *csv.Writer
var ContractExecuteEfficiencyF *csv.Writer
var ConsensusTbftCostF *csv.Writer
var (
	// registerInfo: key: modelName; value: register point info
	registerInfo = make(map[string]string)
)

func ConsensusCliqueCostInit() {
	path := fmt.Sprintf("%s/consensus_clique_cost.csv", Workdir)
	ConsensusCliqueCostF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(ConsensusCliqueCostF, "consensus_clique_cost open failed!")
	}
	defer ConsensusCliqueCostF.Close()
	str := "block_height,clique_start,clique_end,cost_time\n" //需要写入csv的数据，切片类型

	// 写入一条数据，传入数据为切片(追加模式)
	_, err1 := ConsensusCliqueCostF.WriteString(str)
	if err1 != nil {
		log.Println("[consensus_clique_cost] init failed")
	}
	log.Println("[consensus_clique_cost] init succeed")
}

func ContractTimeInit() {
	path := fmt.Sprintf("%s/contract_time.csv", Workdir)
	ContractTimeF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(ContractTimeF, "contract_time open failed!")
	}
	defer ContractTimeF.Close()
	str := "TxHash,ContractAddr,StartTime,EndTime,ExecTime\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := ContractTimeF.WriteString(str)
	if err1 != nil {
		log.Println("[contract_time] init failed")
	}
	log.Println("[contract_time] init succeed")
}

func TransactionPoolInputThroughputInit() {
	path := fmt.Sprintf("%s/transaction_pool_input_throughput.csv", Workdir)
	TransactionPoolInputThroughputF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(TransactionPoolInputThroughputF, "transaction_pool_input_throughput open failed!")
	}
	defer TransactionPoolInputThroughputF.Close()

	str := "measure_time,tx_id,source\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := TransactionPoolInputThroughputF.WriteString(str)
	if err1 != nil {
		log.Println("[transaction_pool_input_throughput] init failed")
	}
	log.Println("[transaction_pool_input_throughput] init succeed")
	// TransactionPoolInputThroughputB = bufio.NewWriterSize(TransactionPoolInputThroughputF, 1024)
	// defer TransactionPoolInputThroughputB.Flush()
}

func NetP2pTransmissionLatencyInit() {
	path := fmt.Sprintf("%s/net_p2p_transmission_latency.csv", Workdir)
	NetP2pTransmissionLatencyF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(NetP2pTransmissionLatencyF, "net_p2p_transmission_latency open failed!")
	}
	defer NetP2pTransmissionLatencyF.Close()
	str := "measure_time,peer_id,peer1_deliver_time,peer2_receive_time,peer2_deliver_time,peer1_receive_time\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := NetP2pTransmissionLatencyF.WriteString(str)
	if err1 != nil {
		log.Println("[net_p2p_transmission_latency] init failed")
	}
	log.Println("[net_p2p_transmission_latency] init succeed")
}

func PeerMessageThroughputInit() {
	path := fmt.Sprintf("%s/peer_message_throughput.csv", Workdir)
	PeerMessageThroughputF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(PeerMessageThroughputF, "peer_message_throughput open failed!")
	}
	defer PeerMessageThroughputF.Close()
	str := "measure_time,message_type,message_size\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := PeerMessageThroughputF.WriteString(str)
	if err1 != nil {
		log.Println("[peer_message_throughput] init failed")
	}
	log.Println("[peer_message_throughput] init succeed")
}

func DbStateWriteRateInit() {
	path := fmt.Sprintf("%s/db_state_write_rate.csv", Workdir)
	DbStateWriteRateF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(DbStateWriteRateF, "db_state_write_rate open failed!")
	}
	defer DbStateWriteRateF.Close()
	str := "measure_time,block_height,block_hash,write_duration\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := DbStateWriteRateF.WriteString(str)
	if err1 != nil {
		log.Println("[db_state_write_rate] init failed")
	}
	log.Println("[db_state_write_rate] init succeed")
}

func DbStateReadRateInit() {
	path := fmt.Sprintf("%s/db_state_read_rate.csv", Workdir)
	DbStateReadRateF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(DbStateReadRateF, "db_state_read_rate open failed!")
	}
	defer DbStateReadRateF.Close()
	str := "measure_time,block_hash,read_duration,\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := DbStateReadRateF.WriteString(str)
	if err1 != nil {
		log.Println("[db_state_read_rate] init failed")
	}
	log.Println("[db_state_read_rate] init succeed")
}

func TxQueueDelayInit() {
	path := fmt.Sprintf("%s/tx_queue_delay.csv", Workdir)
	TxQueueDelayF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(TxQueueDelayF, "tx_queue_delay open failed!")
	}
	defer TxQueueDelayF.Close()
	str := "measure_time,tx_hash,in/outFlag\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := TxQueueDelayF.WriteString(str)
	if err1 != nil {
		log.Println("[tx_queue_delay] init failed")
	}
	log.Println("[tx_queue_delay] init succeed")
}

func BlockCommitDurationStartInit() {
	path := fmt.Sprintf("%s/block_commit_duration_start.csv", Workdir)
	BlockCommitDurationStartF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockCommitDurationStartF, "block_commit_duration_start open failed!")
	}
	defer BlockCommitDurationStartF.Close()
	str := "measure_time,block_height\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockCommitDurationStartF.WriteString(str)
	if err1 != nil {
		log.Println("[block_commit_duration_start] init failed")
	}
	log.Println("[block_commit_duration_start] init succeed")
}

func BlockCommitDurationEndInit() {
	path := fmt.Sprintf("%s/block_commit_duration_end.csv", Workdir)
	BlockCommitDurationEndF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockCommitDurationEndF, "block_commit_duration_end open failed!")
	}
	defer BlockCommitDurationEndF.Close()
	str := "measure_time,block_height,block_hash,block_tx_count\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockCommitDurationEndF.WriteString(str)
	if err1 != nil {
		log.Println("[block_commit_duration_end] init failed")
	}
	log.Println("[block_commit_duration_end] init succeed")
}

func BlockValidationEfficiencyInit() {
	path := fmt.Sprintf("%s/block_validation_efficiency.csv", Workdir)
	BlockValidationEfficiencyF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockValidationEfficiencyF, "block_validation_efficiency open failed!")
	}
	defer BlockValidationEfficiencyF.Close()
	str := "start_time,end_time,exec_time,tx_number\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockValidationEfficiencyF.WriteString(str)
	if err1 != nil {
		log.Println("[block_validation_efficiency] init failed")
	}
	log.Println("[block_validation_efficiency] init succeed")
}

func BlockValidationEfficiencyStartInit() {
	path := fmt.Sprintf("%s/block_validation_efficiency_start.csv", Workdir)
	BlockValidationEfficiencyStartF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockValidationEfficiencyStartF, "block_validation_efficiency_start open failed!")
	}
	defer BlockValidationEfficiencyStartF.Close()
	str := "measure_time,block_hash,block_validation_duration\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockValidationEfficiencyStartF.WriteString(str)
	if err1 != nil {
		log.Println("[block_validation_efficiency_start] init failed")
	}
	log.Println("[block_validation_efficiency_start] init succeed")
}

func BlockValidationEfficiencyEndInit() {
	path := fmt.Sprintf("%s/block_validation_efficiency_end.csv", Workdir)
	BlockValidationEfficiencyEndF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockValidationEfficiencyEndF, "block_validation_efficiency_end open failed!")
	}
	defer BlockValidationEfficiencyEndF.Close()
	str := "measure_time,block_hash,block_tx_count\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockValidationEfficiencyEndF.WriteString(str)
	if err1 != nil {
		log.Println("[block_validation_efficiency_end] init failed")
	}
	log.Println("[block_validation_efficiency_end] init succeed")
}

func TxDelayStartInit() {
	path := fmt.Sprintf("%s/tx_delay_start.csv", Workdir)
	TxDelayStartF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(TxDelayStartF, "tx_delay_start open failed!")
	}
	defer TxDelayStartF.Close()
	str := "measure_time,tx_hash\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := TxDelayStartF.WriteString(str)
	if err1 != nil {
		log.Println("[tx_delay_start] init failed")
	}
	log.Println("[tx_delay_start] init succeed")
}

func TxDelayEndInit() {
	path := fmt.Sprintf("%s/tx_delay_end.csv", Workdir)
	TxDelayEndF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(TxDelayEndF, "tx_delay_end open failed!")
	}
	defer TxDelayEndF.Close()
	str := "measure_time,block_height,tx_hash\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := TxDelayEndF.WriteString(str)
	if err1 != nil {
		log.Println("[tx_delay_end] init failed")
	}
	log.Println("[tx_delay_end] init succeed")

}

func BlockTxConflictRateInit() {
	path := fmt.Sprintf("%s/block_tx_conflict_rate.csv", Workdir)
	BlockTxConflictRateF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockTxConflictRateF, "block_tx_conflict_rate open failed!")
	}
	defer BlockTxConflictRateF.Close()
	str := "measure_time,conflict_count,block_height,block_tx_count\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockTxConflictRateF.WriteString(str)
	if err1 != nil {
		log.Println("[block_tx_conflict_rate] init failed")
	}
	log.Println("[block_tx_conflict_rate] init succeed")
}

func ContractExecuteEfficiencyInit() {
	path := fmt.Sprintf("%s/contract_execute_efficiency.csv", Workdir)
	ContractExecuteEfficiencyF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(ContractExecuteEfficiencyF, "contract_execute_efficiency open failed!")
	}
	defer ContractExecuteEfficiencyF.Close()
	str := "block_height,contract_start_time,contract_end_time,gas\n" //需要写入csv的数据，切片类型

	//写入一条数据，传入数据为切片(追加模式)
	_, err1 := ContractExecuteEfficiencyF.WriteString(str)
	if err1 != nil {
		log.Println("[contract_execute_efficiency] init failed")
	}
	log.Println("[contract_execute_efficiency] init succeed")
}

func ConsensusTbftCostInit() {
	path := fmt.Sprintf("%s/consensus_tbft_cost.csv", Workdir)
	ConsensusTbftCostF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(ConsensusTbftCostF, "consensus_tbft_cost open failed!")
	}
	defer ConsensusTbftCostF.Close()
	str := "ID,Tbfttimestamp,Height,Round,Proposal,Prevote,Precommit,Commit,RoundTotalTime,HeightTotalTime\n" //需要写入csv的数据，切片类型

	// 写入一条数据，传入数据为切片(追加模式)
	_, err1 := ConsensusTbftCostF.WriteString(str)
	if err1 != nil {
		log.Println("[consensus_tbft_cost] init failed")
	}
	log.Println("[consensus_tbft_cost] init succeed")
}

func BlockCommitDurationInit() {
	path := fmt.Sprintf("%s/block_commit_duration.csv", Workdir)
	BlockCommitDurationF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(BlockCommitDurationF, "block_commit_duration open failed!")
	}
	defer BlockCommitDurationF.Close()
	str := "TxHash,BeginTime,EndTime,ExecTime\n" //需要写入csv的数据，切片类型

	// 写入一条数据，传入数据为切片(追加模式)
	_, err1 := BlockCommitDurationF.WriteString(str)
	if err1 != nil {
		log.Println("[block_commit_duration] init failed")
	}
	log.Println("[block_commit_duration] init succeed")
}

func TxinBlockTpsInit() {
	path := fmt.Sprintf("%s/tx_in_block_tps.csv", Workdir)
	TxinBlockTpsF, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(TxinBlockTpsF, "tx_in_block_tps open failed!")
	}
	defer TxinBlockTpsF.Close()
	str := "measure_time,block_height,block_hash,tx_numbert\n" //需要写入csv的数据，切片类型

	// 写入一条数据，传入数据为切片(追加模式)
	_, err1 := TxinBlockTpsF.WriteString(str)
	if err1 != nil {
		log.Println("[tx_in_block_tps] init failed")
	}
	log.Println("[tx_in_block_tps] init succeed")
}

func CreateLog() {
	_, absPath, _, _ := runtime.Caller(0)
	Workdir = filepath.Dir(filepath.Dir(filepath.Dir(absPath))) + "/log"
	_, err := os.ReadDir(Workdir)
	if err != nil {
		err = os.MkdirAll(Workdir, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ConfigInit() {
	CreateLog()
	TransactionPoolInputThroughputInit()
	NetP2pTransmissionLatencyInit()
	PeerMessageThroughputInit()
	DbStateWriteRateInit()
	DbStateReadRateInit()
	ContractTimeInit()
	TxQueueDelayInit()
	BlockCommitDurationStartInit()
	BlockCommitDurationEndInit()
	BlockValidationEfficiencyStartInit()
	BlockValidationEfficiencyEndInit()
	TxDelayStartInit()
	TxDelayEndInit()
	BlockTxConflictRateInit()
	ContractExecuteEfficiencyInit()
	ConsensusTbftCostInit()
	BlockCommitDurationInit()
	TxinBlockTpsInit()
	BlockValidationEfficiencyInit()
	ConsensusCliqueCostInit()
}

func Start(port uint16) error {
	ConfigInit()
	startConfigListener(port)
	return nil
}

func Record(data string, filename string) error {
	// accessLock.RLock()
	allAccess := accessConfig["All"]
	modelAccess := accessConfig[filename]
	// accessLock.RUnlock()
	// 开关的判别取决于两个map，查看配置文件对应的指标是否存在且为true
	if allAccess && modelAccess {
		go safeGoroutine(func() error {
			path := fmt.Sprintf("%s/%s.csv", Workdir, filename)
			file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				return nil
			}
			defer file.Close()

			//写入一条数据，传入数据为切片(追加模式)
			_, err = file.WriteString(data)
			if err != nil {
				return err
			}
			return nil
		}, nil)
	}
	return errors.New("close")
}

// 这里改掉之后峰值从1100到1300
// func Record(data string, filename string) error {
// 	allAccess := accessConfig["All"]
// 	modelAccess := accessConfig[filename]
// 	// accessLock.RUnlock()
// 	// 开关的判别取决于两个map，查看配置文件对应的指标是否存在且为true
// 	if allAccess && modelAccess {
// 		return nil
// 	}
// 	return nil
// }

// GetConfigValue: get the configured value, which can be updated through endpoint [PUT] /config/configvalue
func GetConfigValue(key string) (interface{}, bool) {
	configLock.RLock()
	defer configLock.RUnlock()
	val, ok := configValue[key]
	return val, ok
}
