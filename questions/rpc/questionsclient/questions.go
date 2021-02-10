// Code generated by goctl. DO NOT EDIT!
// Source: questions.proto

//go:generate mockgen -destination ./questions_mock.go -package questionsclient -source $GOFILE

package questionsclient

import (
	"context"

	"datacenter/questions/rpc/questions"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request             = questions.Request
	ConvertResp         = questions.ConvertResp
	ActInfoResp         = questions.ActInfoResp
	AwardInfoResp       = questions.AwardInfoResp
	QuestionsListResp   = questions.QuestionsListResp
	ConvertReq          = questions.ConvertReq
	ActivitiesReq       = questions.ActivitiesReq
	QuestionsAnswerReq  = questions.QuestionsAnswerReq
	QuestionsAnswerResp = questions.QuestionsAnswerResp
	Response            = questions.Response
	AwardListResp       = questions.AwardListResp
	QuestionsResp       = questions.QuestionsResp
	GradeReq            = questions.GradeReq
	TurnTableReq        = questions.TurnTableReq

	Questions interface {
		// 获取 问答抽奖活动信息
		GetActivitiesInfo(ctx context.Context, in *ActivitiesReq) (*ActInfoResp, error)
		// 获取 问答奖品信息
		GetAwardInfo(ctx context.Context, in *ActivitiesReq) (*AwardInfoResp, error)
		// 获取 问答奖品列表
		GetAwardList(ctx context.Context, in *ActivitiesReq) (*AwardListResp, error)
		// 获取 问题列表
		GetQuestionsList(ctx context.Context, in *ActivitiesReq) (*QuestionsListResp, error)
		//  提交 答案
		PostQuestionsChange(ctx context.Context, in *QuestionsAnswerReq) (*QuestionsAnswerResp, error)
		// 获取得分
		GetQuestionsGrade(ctx context.Context, in *GradeReq) (*QuestionsAnswerResp, error)
		// 抽奖
		PostTurnTable(ctx context.Context, in *TurnTableReq) (*AwardInfoResp, error)
		// 填写中奖记录
		PostConvert(ctx context.Context, in *ConvertReq) (*ConvertResp, error)
	}

	defaultQuestions struct {
		cli zrpc.Client
	}
)

func NewQuestions(cli zrpc.Client) Questions {
	return &defaultQuestions{
		cli: cli,
	}
}

// 获取 问答抽奖活动信息
func (m *defaultQuestions) GetActivitiesInfo(ctx context.Context, in *ActivitiesReq) (*ActInfoResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.GetActivitiesInfo(ctx, in)
}

// 获取 问答奖品信息
func (m *defaultQuestions) GetAwardInfo(ctx context.Context, in *ActivitiesReq) (*AwardInfoResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.GetAwardInfo(ctx, in)
}

// 获取 问答奖品列表
func (m *defaultQuestions) GetAwardList(ctx context.Context, in *ActivitiesReq) (*AwardListResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.GetAwardList(ctx, in)
}

// 获取 问题列表
func (m *defaultQuestions) GetQuestionsList(ctx context.Context, in *ActivitiesReq) (*QuestionsListResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.GetQuestionsList(ctx, in)
}

//  提交 答案
func (m *defaultQuestions) PostQuestionsChange(ctx context.Context, in *QuestionsAnswerReq) (*QuestionsAnswerResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.PostQuestionsChange(ctx, in)
}

// 获取得分
func (m *defaultQuestions) GetQuestionsGrade(ctx context.Context, in *GradeReq) (*QuestionsAnswerResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.GetQuestionsGrade(ctx, in)
}

// 抽奖
func (m *defaultQuestions) PostTurnTable(ctx context.Context, in *TurnTableReq) (*AwardInfoResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.PostTurnTable(ctx, in)
}

// 填写中奖记录
func (m *defaultQuestions) PostConvert(ctx context.Context, in *ConvertReq) (*ConvertResp, error) {
	client := questions.NewQuestionsClient(m.cli.Conn())
	return client.PostConvert(ctx, in)
}
