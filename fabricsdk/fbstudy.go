/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */



package main

import (
	//"os"
	//"path"
	//"testing"
	//"time"

	//"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	//ca "github.com/hyperledger/fabric-sdk-go/api/apifabca"
	//fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	//"github.com/hyperledger/fabric-sdk-go/api/apitxn"
	//pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

	fabricapi "github.com/hyperledger/fabric-sdk-go/def/fabapi"
	//"github.com/hyperledger/fabric-sdk-go/pkg/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/errors"
	//"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/events"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/orderer"
	//admin "github.com/hyperledger/fabric-sdk-go/pkg/fabric-txn/admin"
	//"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"fmt"
	identityImpl "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/identity"
	//hfc "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	//"github.com/cloudflare/cfssl/api/client
	//config "github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	//fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"

	"io/ioutil"
	//"github.com/cloudflare/cfssl/api/client"
	"encoding/hex"

	afc "github.com/hyperledger/fabric-sdk-go/api/apifabclient"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn"

	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"

	peerapi "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/peer"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/events"

	//ca "github.com/hyperledger/fabric-sdk-go/api/apifabca"
	/*fabricCAClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-ca-client"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"


	bccspFactory "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/bccsp/factory"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"*/


)


/*type fbconfig struct {
	tlsEnabled bool
	errorCase  bool
}*/

func main() {

	//fabric_local()
	test()
}



/*
func fabric_ca(){

	configImpl, err := config.InitConfig("./fabricsdk/config_test.yaml")
	if err != nil {
		check( err ,"finderror" )
	}

	caConfig, err := configImpl.CAConfig("Org1")
	if err != nil {
		check( err ,"finderror" )
	}

	caname :=caConfig.CAName

	fmt.Println(" %s ", caname)

	client := hfc.NewClient(configImpl)


	err = bccspFactory.InitFactories(configImpl.CSPConfig())
	if err != nil {
		check( err ,"finderror" )
	}

	cryptoSuite := bccspFactory.GetDefault()
	client.SetCryptoSuite(cryptoSuite)
	stateStore, err := kvs.CreateNewFileKeyValueStore("/tmp/enroll_user")

	client.SetStateStore(stateStore)


	caClient, err := fabricCAClient.NewFabricCAClient(configImpl, "Org1")
	if err != nil {
		check( err ,"caClient error " )
	}

	adminUser, err := client.LoadUserFromStateStore("admin")

	if err != nil {
		check( err ,"finderror" )
	}

	if adminUser == nil {

		key, cert, err := caClient.Enroll("admin", "adminpw")
		if err != nil {
			check( err ," Enroll return error: %v " )
		}
		if key == nil {
			check( err ," private key return from Enroll is nil " )
		}
		if cert == nil {
			check( err ," cert return from Enroll is nil " )
		}

	}

}
*/

func test(){

	//读取配置文件
	sdkOptions := fabricapi.Options{
		ConfigFile: "./fabricsdk/config_test.yaml",
	}

	//创建SDK代理
	sdk, _ := fabricapi.NewSDK(sdkOptions)
	session, _ := sdk.NewPreEnrolledUserSession("org1", "Admin")

	//创建Golang的 fabric客户端代理
	client, _ := sdk.NewSystemClient(session)

    //创建通道代理，通道名为:roberttestchannel12
	channel, _ := client.NewChannel("roberttestchannel12")

	//创建Orderer节点代理
	orderer, _ := orderer.NewOrderer("grpc://192.168.23.212:7050", "", "", client.Config())
	channel.AddOrderer(orderer)

	//创建Peer节点代理
	peer ,_ := fabricapi.NewPeer("grpc://192.168.23.212:7051","","",client.Config())
	channel.AddPeer(peer)

	//获取当前通道的信息
	blockchainInfo, _ := channel.QueryInfo()
	fmt.Println(" the peer block height  %d",blockchainInfo.Height)



	//获取当前通道的详细信息
	/*bciAfterTx, err4 := channel.QueryInfo()
	check(err4,"_6")
	fmt.Println(" the peer block height  %d",bciAfterTx.Height)*/

}


func fabric_local(){

		sdkOptions := fabricapi.Options{
			ConfigFile: "./fabricsdk/config_test.yaml",
		}

		sdk, err := fabricapi.NewSDK(sdkOptions)
		if err != nil {
			check(err, "SDK init failed")
		}

		session, err := sdk.NewPreEnrolledUserSession("org1", "Admin")
		if err != nil {
			check(err, "failed getting admin user session for org")
		}

		sc, err := sdk.NewSystemClient(session)
		if err != nil {
			check( err ,"NewSystemClient failed")
		}

		client := sc

		channel, err := client.NewChannel("roberttestchannel12")

		if err != nil {
			 check( err , "NewChannel failed")
		}

		orderer, err := orderer.NewOrderer("grpc://192.168.23.212:7050", "", "", client.Config())
		if err != nil {
			check(err, "NewOrderer failed")
		}
		channel.AddOrderer(orderer)


		peer ,err := fabricapi.NewPeer("grpc://192.168.23.212:7051","","",client.Config())
		if err != nil {
			check(err, "NewOrderer failed")
		}
		channel.AddPeer(peer)

		peer188 ,err := fabricapi.NewPeer("grpc://172.16.10.188:7051","","",client.Config())
		if err != nil {
			check(err, "NewPeer  failed")
		}
		channel.AddPeer(peer188)




		//获取当前通道的详细信息
		/*bciAfterTx, err4 := channel.QueryInfo()
		check(err4,"_6")
		fmt.Println(" the peer block height  %d",bciAfterTx.Height)*/


		//查询Peer节点加入的所有通道
		/*channels,err := client.QueryChannels(peer)

		for _ , responsechannel := range channels.Channels{
			fmt.Println(" the channel info is : %s " , responsechannel.ChannelId  )
		}*/


		//根据区块编号获取区块信息
		/*block, err := channel.QueryBlock(23)
		fmt.Println(" The block info : %s  " , hex.EncodeToString(block.Header.PreviousHash))*/


		//根据区块链HASH获取区块详细信息
		/*let blockinfobyhash = yield channel.queryBlockByHash(new Buffer("ec298dc1cd1f0e0a3f6d6e25b5796e7b5e4d668aeb6ec3a90b4aa6bb1a7f0c17","hex"),peer)
		console.info(  JSON.stringify(blockinfobyhash ) )*/

		//blockinfo , err := channel.QueryBlockByHash(block.Header.PreviousHash)
		/*blockhash , err :=  hex.DecodeString("7376086e18a8ddbc40d318557c39987fd538c64340aa0df191e1062b935e147e")
		blockinfo , err := channel.QueryBlockByHash( blockhash )
		fmt.Println(" The block info : %s  " , blockinfo.String())*/

		//查询已经install的chaincode
		/*installchaincodes , err := client.QueryInstalledChaincodes(peer)
		for _ , responseinstall := range installchaincodes.Chaincodes{

			fmt.Println(" chaincode info is  : %s   %s " , responseinstall.Version , responseinstall.Path )

		}*/

		// 查询已经实例化的Chaincode
		/*channel.SetPrimaryPeer(peer)
		instantChaincodes, err := channel.QueryInstantiatedChaincodes();
		for _ , responseinstant := range instantChaincodes.Chaincodes{

			fmt.Println(" chaincode info is  : %s   %s " , responseinstant.Version , responseinstant.Path )
		}*/



		//根据交易编号获取交易详细信息

		/* let resulttxinfo = yield channel.queryTransaction("56f51f9a54fb4755fd68c6c24931234a59340f7c98308374e9991d276d7d4a96", peer);
		 console.info(  JSON.stringify( resulttxinfo ) )
		*/

		/*tran , err := channel.QueryTransaction("56f51f9a54fb4755fd68c6c24931234a59340f7c98308374e9991d276d7d4a96")
		fmt.Println(" transcaion info is :  %s  " , tran.String())
		*/


		//发起交易(查询交易)

		/*targets := peerapi.PeersToTxnProcessors(channel.Peers())
		client.SetUserContext(session.Identity())
		request := apitxn.ChaincodeInvokeRequest{
			Targets:     targets,
			ChaincodeID: "cc_endfinlshed",
			Fcn:         "query",
			Args:        [][]byte{[]byte("a")},
		}

		queryResponses, err := channel.QueryByChaincode(request)
		if err != nil {
			check(err,"QueryByChaincode failed %s")
		}


		for _ , parmbytes := range queryResponses{
			fmt.Println(" chaincode query info is :  %s  " , string(parmbytes))
		}*/



		// 发起交易，写入交易

		targets := peerapi.PeersToTxnProcessors(channel.Peers())
		transientData := make(map[string][]byte)

		request := apitxn.ChaincodeInvokeRequest{
			Targets:      targets,
			Fcn:          "invoke",
			Args:         [][]byte{[]byte("a"),[]byte("b"),[]byte("1")},
			TransientMap: transientData,
			ChaincodeID:  "cc_endfinlshed",
		}

		transactionProposalResponses, txnID, err := channel.SendTransactionProposal(request)
		fmt.Println("   tx id  : %s",txnID)
		if err != nil {
			check(err," send transtion error ")
		}

		for _, v := range transactionProposalResponses {

			if v.Err != nil {
				check(v.Err, "endorser %s failed")
			}

		}

		tx, err := channel.CreateTransaction(transactionProposalResponses)
		transactionResponse, err := channel.SendTransaction(tx)


		fmt.Println(" srarch result   %s " , transactionResponse )


		//eventHub, err := events.NewEventHub(client)

}


// getEventHub initilizes the event hub
func getEventHub( client fab.FabricClient) (fab.EventHub, error) {

	eventHub, err := events.NewEventHub(client)

	if err != nil {
		return nil, errors.WithMessage(err, "NewEventHub failed")
	}

	foundEventHub := false
	peerConfig, err := client.Config().PeersConfig("org1")

	if err != nil {
		return nil, errors.WithMessage(err, "PeersConfig failed")
	}
	for _, p := range peerConfig {
		if p.URL != "" {

			serverHostOverride := ""
			if str, ok := p.GRPCOptions["ssl-target-name-override"].(string); ok {
				serverHostOverride = str
			}
			eventHub.SetPeerAddr(p.EventURL, p.TLSCACerts.Path, serverHostOverride)
			foundEventHub = true
			break
		}
	}

	if !foundEventHub {
		return nil, errors.New("event hub configuration not found")
	}

	return eventHub, nil
}


/**
  创建交提案
 */
func  CreateAndSendTransactionProposal(channel afc.Channel,
	                                   chainCodeID string,
	                                   fcn string,
	                                   	args [][]byte,
                                   		targets []apitxn.ProposalProcessor,
                              			transientData map[string][]byte) ([]*apitxn.TransactionProposalResponse, apitxn.TransactionID, error) {

	request := apitxn.ChaincodeInvokeRequest{
		Targets:      targets,
		Fcn:          fcn,
		Args:         args,
		TransientMap: transientData,
		ChaincodeID:  chainCodeID,
	}
	transactionProposalResponses, txnID, err := channel.SendTransactionProposal(request)
	if err != nil {
		return nil, txnID, err
	}

	for _, v := range transactionProposalResponses {
		if v.Err != nil {
			return nil, txnID, errors.Wrapf(v.Err, "endorser %s failed", v.Endorser)
		}
	}

	return transactionProposalResponses, txnID, nil
}




func temp(){


	//fmt.Print("ddd")

	sdkOptions := fabricapi.Options{
		ConfigFile: "./fabricsdk/config_test.yaml",
	}


	sdk, err := fabricapi.NewSDK(sdkOptions)

	if err != nil {
		fmt.Println( errors.WithMessage(err, "SDK init failed"))
	}


	//session, err := sdk.NewPreEnrolledUserSession("Org1MSP" ,"Admin")

	countext,err := sdk.NewContext("Org1MSP")
	user := identityImpl.NewUser("Admin","Org1MSP");


	cer, err := ioutil.ReadFile("/project/opt_fabric/fabricconfig/crypto-config/peerOrganizations/org1.robertfabrictest.com/users/Admin@org1.robertfabrictest.com/msp/signcerts/Admin@org1.robertfabrictest.com-cert.pem")
	check(err,"1")

	privatekey , err1 := ioutil.ReadFile("/project/opt_fabric/fabricconfig/crypto-config/peerOrganizations/org1.robertfabrictest.com/users/Admin@org1.robertfabrictest.com/msp/keystore/b031338f76290f089d330b064d4534202a49ae8d65ca5d266c377bc46812a884_sk")
	check(err1,"2")

	pfstring := hex.EncodeToString(privatekey)

	fmt.Println("dddd   + " + pfstring )


	pk,err2 := sdk.CryptoSuiteProvider().GetKey(privatekey)
	check(err2,"2.1")

	user.SetEnrollmentCertificate(cer)
	user.SetPrivateKey(pk)

	session, err := sdk.NewSession(countext,user)

	if err != nil {
		fmt.Println( errors.WithMessage(err, "failed getting admin user session for org"))
	}

	sc, err := sdk.NewSystemClient(session)


	if err != nil {
		fmt.Println( errors.WithMessage(err, "NewSystemClient failed") )
	}



	channel, err_1 := sc.NewChannel("roberttestchannel" )
	check(err_1,"_1")

	orderer, err_2 := orderer.NewOrderer("grpc://192.168.23.212:7050", "", "", sc.Config())
	check(err_2,"_2")

	newpeer, err_3 := fabricapi.NewPeer("grpc://192.168.23.212:7051", "", "", sc.Config())
	check(err_3,"_3")


	err_4 :=channel.AddOrderer(orderer)
	check(err_4,"_4")

	channel.AddPeer(newpeer)

	channel.SetPrimaryPeer(newpeer)

	primaryPeer := channel.PrimaryPeer()
	_, err_5 := sc.QueryChannels(primaryPeer)
	check(err_5,"_5")

	/*for _, responseChannel := range response.Channels {
		if responseChannel.ChannelId == channel.Name() {

		}
	}*/

	bciAfterTx, err4 := channel.QueryInfo()

	check(err4,"_6")

	fmt.Println(" the peer block height  %d",bciAfterTx.Height)



}



func check(e error , num string) {
	if e != nil {
		fmt.Println( errors.WithMessage(e, " find a error "+num))
	}
}