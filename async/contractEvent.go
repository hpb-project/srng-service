package async

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hpb-project/srng-service/models"
	"math/big"
	"strings"
)

func HorseArenaContractHandler(vLog types.Log) error {
	logs.Info("handler arena contract logs")
	arenaContractAbi := horseArenaContract.HorseArenaContractABI
	abiInfo, _ := abi.JSON(strings.NewReader(arenaContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.MintArena:
			rp, err := abiInfo.Unpack("MintArena", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("MintArena", rp)
			mintArena := models.MintArena{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Types:       rp[3].(*big.Int).String(),
				Mort_Amount: rp[4].(*big.Int).String(),
				Name:        rp[5].(string),
				Status:      rp[6].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "mint_arena",
			}
			err = models.Insert(mintArena)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(mintArena))
			}
		case common.CloseArena:
			rp, err := abiInfo.Unpack("CloseArena", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CloseArena", rp)
			closeArena := models.CloseArena{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Mort_Amount: rp[3].(*big.Int).String(),
				Status:      rp[4].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "close_arena",
			}
			err = models.Insert(closeArena)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(closeArena))
			}
		case common.ArenaTransfer:
			logs.Info("got event arenaTransfer")
			rp, err := abiInfo.Unpack("ArenaTransfer", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("ArenaTransfer", rp)
			arenaTransfer := models.ArenaTransfer{
				TxFrom:      rp[0].(comm.Address).String(),
				TxTo:        rp[1].(comm.Address).String(),
				ArenaId:     rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "arenaTransfer",
			}
			err = models.Insert(arenaTransfer)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(arenaTransfer))
			}
		}
	}
	return nil
}

func HorseArenaExtraHandler(vLog types.Log) error {
	logs.Info("handler arena extra logs.")
	arenaExtraContractAbi := horseArenaExtra.HorseArenaExtraABI
	abiInfo, _ := abi.JSON(strings.NewReader(arenaExtraContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.EndGame:
			rp, err := abiInfo.Unpack("EndGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("EndGame", rp)
			endGame := models.EndGame{
				Game_Id:     rp[0].(*big.Int).String(),
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "end_game",
			}
			err = models.Insert(endGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(endGame))
			}
		case common.SetHorseIntegral:
			rp, err := abiInfo.Unpack("SetHorseIntegral", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("SetHorseIntegral", rp)
			setHorseIntegral := models.SetHorseIntegral{
				Horse_Id:      rp[0].(*big.Int).String(),
				Count:         rp[1].(*big.Int).String(),
				Time:          TimestampToMsillisecond(rp[2].(*big.Int).String()),
				IntegralYear:  rp[3].(*big.Int).String(),
				IntegralMonth: rp[4].(*big.Int).String(),
				IntegralWeek:  rp[5].(*big.Int).String(),
				Event_Index:   vLog.Index,
				Tx_Hash:       vLog.TxHash.String(),
				Event_Name:    "set_horse_integral",
			}
			err = models.Insert(setHorseIntegral)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(setHorseIntegral))
			}
		case common.SetHorseGradeSc:
			rp, err := abiInfo.Unpack("SetHorseGradeSc", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("SetHorseGradeSc", rp)
			setHorseGradeSc := models.SetHorseGradeSc{
				Horse_Id:    rp[0].(*big.Int).String(),
				Count:       rp[1].(*big.Int).String(),
				Mark:        rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "set_horse_grade_sc",
			}
			err = models.Insert(setHorseGradeSc)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(setHorseGradeSc))
			}
		case common.EndGameOfHorse:
			rp, err := abiInfo.Unpack("EndGameOfHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("EndGameOfHorse", rp)
			endGameOfHorse := models.EndGameOfHorse{
				Horse_Id:    rp[0].(*big.Int).String(),
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Arena_Id:    rp[3].(*big.Int).String(),
				Game_Id:     rp[4].(*big.Int).String(),
				Race_Type:   rp[5].(*big.Int).String(),
				Distance:    rp[6].(*big.Int).String(),
				Energy:      rp[7].(*big.Int).String(),
				Race_Count:  rp[8].(*big.Int).String(),
				Win_Count:   rp[9].(*big.Int).String(),
				Grade:       rp[10].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "end_game_of_horse",
			}
			err = models.Insert(endGameOfHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(endGameOfHorse))
			}
		}
	}
	return nil
}

func HorseArenaExtra2Handler(vLog types.Log) error {
	logs.Info("handler horse arena extra2 logs.")
	arenaContractAbi := horseArenaExtra2.HorseArenaExtra2ABI
	abiInfo, _ := abi.JSON(strings.NewReader(arenaContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.ApplyGame:
			rp, err := abiInfo.Unpack("ApplyGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("ApplyGame", rp)
			applyGame := models.ApplyGame{
				Account:     rp[0].(comm.Address).String(),
				Horse_Id:    rp[1].(*big.Int).String(),
				Arena_Id:    rp[2].(*big.Int).String(),
				Race_Type:   rp[3].(*big.Int).String(),
				Distance:    rp[4].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[5].(*big.Int).String()),
				Status:      rp[6].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "apply_game",
			}
			err = models.Insert(applyGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(applyGame))
			}
		case common.CancelApplyGame:
			rp, err := abiInfo.Unpack("CancelApplyGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelApplyGame", rp)
			cancelApplyGame := models.CancelApplyGame{
				Account:     rp[0].(comm.Address).String(),
				Horse_Id:    rp[1].(*big.Int).String(),
				Status:      rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Arena_Id:    rp[4].(*big.Int).String(),
				Game_Id:     rp[5].(*big.Int).String(),
				Race_Type:   rp[6].(*big.Int).String(),
				Distance:    rp[7].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_apply_game",
			}
			err = models.Insert(cancelApplyGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelApplyGame))
			}
		}
	}
	return nil
}
func HorseArenaExtra3Handler(vLog types.Log) error {
	logs.Info("handler arena extra3 logs.")
	arenaContractAbi := horseArenaExtra3.HorseArenaExtra3ABI
	abiInfo, _ := abi.JSON(strings.NewReader(arenaContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.StartGame:
			rp, err := abiInfo.Unpack("StartGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("StartGame", rp)
			var t1 []string
			for _, v := range rp[0].([]*big.Int) {
				t1 = append(t1, v.String())
			}
			startGame := models.StartGame{
				Horse_Ids:   t1,
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				RaceId:      rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "start_game",
			}
			err = models.Insert(startGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(startGame))
			}
		case common.StartGameOfRace:
			rp, err := abiInfo.Unpack("StartGameOfRace", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("StartGameOfRace", rp)
			startGameOfRace := models.StartGameOfRace{
				Game_Id:     rp[0].(*big.Int).String(),
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "start_game_of_race",
			}
			err = models.Insert(startGameOfRace)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(startGameOfRace))
			}
		case common.CancelGame:
			rp, err := abiInfo.Unpack("CancelGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelGame", rp)
			var t1 []string
			for _, v := range rp[0].([]*big.Int) {
				t1 = append(t1, v.String())
			}
			cancelGame := models.CancelGame{
				Game_Id:     t1,
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_game",
			}
			err = models.Insert(cancelGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelGame))
			}
		case common.CancelGameOfHorse:
			rp, err := abiInfo.Unpack("CancelGameOfHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelGameOfHorse", rp)
			var t1 []string
			for _, v := range rp[0].([]*big.Int) {
				t1 = append(t1, v.String())
			}
			cancelGameOfHorse := models.CancelGameOfHorse{
				Horse_Ids:   t1,
				Status:      rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_game_of_horse",
			}
			err = models.Insert(cancelGameOfHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelGameOfHorse))
			}
		case common.CancelGameOfArena:
			rp, err := abiInfo.Unpack("CancelGameOfArena", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelGameOfArena", rp)
			var t1 []string
			for _, v := range rp[2].([]*big.Int) {
				t1 = append(t1, v.String())
			}
			cancelGameOfArena := models.CancelGameOfArena{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Game_Id:     t1,
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Count:       rp[4].(*big.Int).String(),
				Total_Count: rp[5].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_game_of_arena",
			}
			err = models.Insert(cancelGameOfArena)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelGameOfArena))
			}
		case common.UptArena:
			rp, err := abiInfo.Unpack("UptArena", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UptArena", rp)
			uptArena := models.UptArena{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Game_Id:     rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Count:       rp[4].(*big.Int).String(),
				Total_Count: rp[5].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "upt_arena",
			}
			err = models.Insert(uptArena)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(uptArena))
			}
		case common.CreateGame:
			rp, err := abiInfo.Unpack("CreateGame", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CreateGame", rp)
			createGame := models.CreateGame{
				Account:     rp[0].(comm.Address).String(),
				Game_Id:     rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Level:       rp[3].(*big.Int).String(),
				Race_Type:   rp[4].(*big.Int).String(),
				Distance:    rp[5].(*big.Int).String(),
				Status:      rp[6].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "create_game",
			}
			err = models.Insert(createGame)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(createGame))
			}
		}
	}
	return nil
}

func HorseEquipContractHandler(vLog types.Log) error {
	logs.Info("handler horse equip contract logs.")
	horseEquipContractAbi := horseEquipContract.HorseEquipContractABI
	abiInfo, _ := abi.JSON(strings.NewReader(horseEquipContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.MintEquip:
			rp, err := abiInfo.Unpack("MintEquip", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("MintEquip", rp)
			mintEquip := models.MintEquip{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				To:          rp[3].(comm.Address).String(),
				Types:       rp[4].(*big.Int).String(),
				Style:       rp[5].(*big.Int).String(),
				Status:      rp[6].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "mint_equip",
			}
			err = models.Insert(mintEquip)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(mintEquip))
			}
		case common.Sell:
			rp, err := abiInfo.Unpack("Sell", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("Sell", rp)
			sell := models.Sell{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Kind:        rp[2].(*big.Int).String(),
				Coin:        rp[3].(*big.Int).String(),
				Price:       rp[4].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[5].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "sell",
			}
			err = models.Insert(sell)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(sell))
			}
		case common.CancelSell:
			rp, err := abiInfo.Unpack("CancelSell", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelSell", rp)
			cancelSell := models.CancelSell{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Status:      rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_sell",
			}
			err = models.Insert(cancelSell)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelSell))
			}
		case common.Buy:
			rp, err := abiInfo.Unpack("Buy", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("Buy", rp)
			buy := models.Buy{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Status:      rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "buy",
			}
			err = models.Insert(buy)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(buy))
			}
		case common.UnloadEquip:
			rp, err := abiInfo.Unpack("UnloadEquip", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UnloadEquip", rp)
			unloadEquip := models.UnloadEquip{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Status:      rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "unload_equip",
			}
			err = models.Insert(unloadEquip)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(unloadEquip))
			}
		case common.UnloadEquipOfHorse:
			rp, err := abiInfo.Unpack("UnloadEquipOfHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UnloadEquipOfHorse", rp)
			unloadEquipOfHorse := models.UnloadEquipOfHorse{
				Account:     rp[0].(comm.Address).String(),
				Horse_Id:    rp[1].(*big.Int).String(),
				Types:       rp[2].(*big.Int).String(),
				Status:      rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "unload_equip_of_horse",
			}
			err = models.Insert(unloadEquipOfHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(unloadEquipOfHorse))
			}
		case common.EquipTransfer:
			logs.Info("got event equipTransfer")
			rp, err := abiInfo.Unpack("EquipTransfer", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("EquipTransfer", rp)
			equipTransfer := models.EquipTransfer{
				TxFrom:      rp[0].(comm.Address).String(),
				TxTo:        rp[1].(comm.Address).String(),
				EquipId:     rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "equipTransfer",
			}
			err = models.Insert(equipTransfer)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(equipTransfer))
			}
		}

	}
	return nil
}

func HorseRaceContractHandler(vLog types.Log) error {
	logs.Info("handler horse race contract logs.")
	horseRaceContractAbi := horseRaceContract.HorseRaceContractABI
	abiInfo, _ := abi.JSON(strings.NewReader(horseRaceContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.Breeding:
			rp, err := abiInfo.Unpack("Breeding", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("Breeding", rp)
			breeding := models.Breeding{
				Account:         rp[0].(comm.Address).String(),
				Token_Id:        rp[1].(*big.Int).String(),
				StallId:         rp[2].(*big.Int).String(),
				NewHorse_Id:     rp[3].(*big.Int).String(),
				Name:            rp[4].(string),
				GenerationScore: rp[5].(*big.Int).String(),
				Gender:          rp[6].(*big.Int).String(),
				IntegralTime:    TimestampToMsillisecond(rp[7].(*big.Int).String()),
				EnergyTime:      TimestampToMsillisecond(rp[8].(*big.Int).String()),
				Status:          rp[9].(*big.Int).String(),
				Event_Index:     vLog.Index,
				Tx_Hash:         vLog.TxHash.String(),
				Event_Name:      "breeding",
			}
			err = models.Insert(breeding)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(breeding))
			}
		case common.Breeding1:
			rp, err := abiInfo.Unpack("Breeding1", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("Breeding1", rp)
			breeding1 := models.Breeding1{
				NewHorse_Id: rp[0].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[1].(*big.Int).String()),
				Color:       rp[2].(string),
				Gene:        rp[3].(string),
				M_Gene:      rp[4].(*big.Int).String(),
				S_Gene:      rp[5].(*big.Int).String(),
				Tra_Value:   rp[6].(*big.Int).String(),
				Energy:      rp[7].(*big.Int).String(),
				Grade:       rp[8].(*big.Int).String(),
				Integral:    rp[9].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "breeding1",
			}
			err = models.Insert(breeding1)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(breeding1))
			}
		case common.BreedingOfHorse:
			rp, err := abiInfo.Unpack("BreedingOfHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("BreedingOfHorse", rp)
			breedingOfHorse := models.BreedingOfHorse{
				Horse_Id:    rp[0].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[1].(*big.Int).String()),
				Count:       rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "breeding_of_horse",
			}
			err = models.Insert(breedingOfHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(breedingOfHorse))
			}
		}
	}
	return nil
}

func HorseRaceExtra1Handler(vLog types.Log) error {
	logs.Info("handler horse race extra1 logs.")
	horseRaceExtra1ContractAbi := horseRaceExtra1.HorseRaceExtra1ABI
	abiInfo, _ := abi.JSON(strings.NewReader(horseRaceExtra1ContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.TrainingHorses:
			rp, err := abiInfo.Unpack("TrainingHorses", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("TrainingHorses", rp)
			trainingHorses := models.TrainingHorses{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Value:       rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "training_horses",
			}
			err = models.Insert(trainingHorses)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(trainingHorses))
			}
		case common.UptHorseName:
			rp, err := abiInfo.Unpack("UptHorseName", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UptHorseName", rp)
			uptHorseName := models.UptHorseName{
				Token_Id:    rp[0].(*big.Int).String(),
				Name:        rp[1].(string),
				Count:       rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "upt_horse_name",
			}
			err = models.Insert(uptHorseName)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(uptHorseName))
			}
		case common.HorseDeco:
			rp, err := abiInfo.Unpack("HorseDeco", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("HorseDeco", rp)
			horseDeco := models.HorseDeco{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Types:       rp[2].(*big.Int).String(),
				Equip_Id:    rp[3].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[4].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "horse_deco",
			}
			err = models.Insert(horseDeco)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(horseDeco))
			}
		case common.HorseDecoOfEquip:
			rp, err := abiInfo.Unpack("HorseDecoOfEquip", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("HorseDecoOfEquip", rp)
			horseDecoOfEquip := models.HorseDecoOfEquip{
				Equip_Id:    rp[0].(*big.Int).String(),
				Status:      rp[1].(*big.Int).String(),
				Token_Id:    rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "horse_deco_of_equip",
			}
			err = models.Insert(horseDecoOfEquip)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(horseDecoOfEquip))
			}
		case common.SetHorseGene:
			rp, err := abiInfo.Unpack("SetHorseGene", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("SetHorseGene", rp)
			setHorseGene := models.SetHorseGene{
				Token_Id:     rp[0].(*big.Int).String(),
				Grip_Gene:    rp[1].(string),
				Acc_Gene:     rp[2].(string),
				End_Gene:     rp[3].(string),
				Speed_Gene:   rp[4].(string),
				Turn_To_Gene: rp[5].(string),
				Control_Gene: rp[6].(string),
				Event_Index:  vLog.Index,
				Tx_Hash:      vLog.TxHash.String(),
				Event_Name:   "set_horse_gene",
			}
			err = models.Insert(setHorseGene)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(setHorseGene))
			}
		case common.InitHorseGrade:
			logs.Info("got event initHorseGrade")
			rp, err := abiInfo.Unpack("InitHorseGrade", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("InitHorseGrade", rp)
			initHorseGrade := models.InitHorseGrade{
				Horse_Id:    rp[0].(*big.Int).String(),
				Grade:       rp[1].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "init_horse_grade",
			}
			err = models.Insert(initHorseGrade)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(initHorseGrade))
			}
		case common.UnloadEquip:
			rp, err := abiInfo.Unpack("UnloadEquip", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UnloadEquip", rp)
			unloadEquip := models.UnloadEquip{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Status:      rp[2].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "unload_equip",
			}
			err = models.Insert(unloadEquip)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(unloadEquip))
			}
		case common.UnloadEquipOfHorse:
			rp, err := abiInfo.Unpack("UnloadEquipOfHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("UnloadEquipOfHorse", rp)
			unloadEquipOfHorse := models.UnloadEquipOfHorse{
				Account:     rp[0].(comm.Address).String(),
				Horse_Id:    rp[1].(*big.Int).String(),
				Types:       rp[2].(*big.Int).String(),
				Status:      rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "unload_equip_of_horse",
			}
			err = models.Insert(unloadEquipOfHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(unloadEquipOfHorse))
			}
		case common.HorseTransfer:
			logs.Info("got event horseTransfer")
			rp, err := abiInfo.Unpack("HorseTransfer", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("HorseTransfer", rp)
			horseTransfer := models.HorseTransfer{
				TxFrom:      rp[0].(comm.Address).String(),
				TxTo:        rp[1].(comm.Address).String(),
				HorseId:     rp[2].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[3].(*big.Int).String()),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "horseTransfer",
			}
			err = models.Insert(horseTransfer)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(horseTransfer))
			}
		}
	}
	return nil
}

func HorseRaceExtra2Handler(vLog types.Log) error {
	logs.Info("handle horse race extra2 logs.")
	horseRaceExtra2ContractAbi := horseRaceExtra2.HorseRaceExtra2ABI
	abiInfo, _ := abi.JSON(strings.NewReader(horseRaceExtra2ContractAbi))
	{
		method := vLog.Topics[0]
		switch method.String() {
		case common.SellHorse:
			rp, err := abiInfo.Unpack("SellHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("SellHorse", rp)
			sellHorse := models.SellHorse{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Coin:        rp[2].(*big.Int).String(),
				Price:       rp[3].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[4].(*big.Int).String()),
				Status:      rp[5].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "sell_horse",
			}
			err = models.Insert(sellHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(sellHorse))
			}
		case common.SireHorse:
			rp, err := abiInfo.Unpack("SireHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("SireHorse", rp)
			sellHorse := models.SireHorse{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Coin:        rp[2].(*big.Int).String(),
				Price:       rp[3].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[4].(*big.Int).String()),
				Status:      rp[5].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "sire_horse",
			}
			err = models.Insert(sellHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(sellHorse))
			}
		case common.CancelSellHorse:
			rp, err := abiInfo.Unpack("CancelSellHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelSellHorse", rp)
			cancelSellHorse := models.CancelSellHorse{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Status:      rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_sell_horse",
			}
			err = models.Insert(cancelSellHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelSellHorse))
			}
		case common.CancelSireHorse:
			rp, err := abiInfo.Unpack("CancelSireHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("CancelSireHorse", rp)
			cancelSireHorse := models.CancelSireHorse{
				Account:     rp[0].(comm.Address).String(),
				Token_Id:    rp[1].(*big.Int).String(),
				Time:        TimestampToMsillisecond(rp[2].(*big.Int).String()),
				Status:      rp[3].(*big.Int).String(),
				Event_Index: vLog.Index,
				Tx_Hash:     vLog.TxHash.String(),
				Event_Name:  "cancel_sire_horse",
			}
			err = models.Insert(cancelSireHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(cancelSireHorse))
			}
		case common.BuyHorse:
			rp, err := abiInfo.Unpack("BuyHorse", vLog.Data)
			if err != nil {
				panic(err)
			}
			logs.Info("BuyHorse", rp)
			buyHorse := models.BuyHorse{
				Account:           rp[0].(comm.Address).String(),
				Token_Id:          rp[1].(*big.Int).String(),
				Coin:              rp[2].(*big.Int).String(),
				Price:             rp[3].(*big.Int).String(),
				Status:            rp[4].(*big.Int).String(),
				Modify_Name_Times: rp[5].(*big.Int).String(),
				Event_Index:       vLog.Index,
				Tx_Hash:           vLog.TxHash.String(),
				Event_Name:        "buy_horse",
			}
			err = models.Insert(buyHorse)
			if err != nil {
				cache.Redis.LpushByte(common.EventList, models.Marshal(buyHorse))
			}
		}
	}
	return nil
}

func TimestampToMsillisecond(second string) string {
	if len(second) > 0 {
		return second + "000"
	}
	return ""
}
