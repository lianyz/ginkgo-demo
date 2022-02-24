package ginkgo_demo_test

import (
	gopher "github.com/lianyz/ginkgo-demo"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func mockInputData() ([]gopher.Gopher, error) {
	inputData := []gopher.Gopher{
		{
			Name:   "菜刀",
			Gender: "男",
			Age:    18,
		},
		{
			Name:   "小西瓜",
			Gender: "女",
			Age:    19,
		},
		{
			Name:   "机器铃砍菜刀",
			Gender: "男",
			Age:    17,
		},
		{
			Name:   "小菜刀",
			Gender: "男",
			Age:    20,
		},
	}

	return inputData, nil
}

var _ = Describe("Gopher", func() {

	BeforeEach(func() {
		By("当测试不通过时，我会在这里打印一个消息 【BeforeEach】")
	})

	inputData, err := mockInputData()

	Describe("校验输入数据", func() {

		Context("当获取数据没有错误发生时", func() {
			It("它应该时接收数据成功了的", func() {
				gomega.Expect(err).Should(gomega.BeNil())
			})
		})

		Context("当获取的数据校验失败时", func() {
			It("当数据校验返回错误为：名字太短，不能小于3 时", func() {
				gomega.Expect(gopher.Validate(inputData[0])).
					Should(gomega.MatchError("名字太短，不能小于3"))
			})

			It("当数据校验返回错误为：只要男的 时", func() {
				gomega.Expect(gopher.Validate(inputData[1])).
					Should(gomega.MatchError("只要男的"))
			})

			It("当数据校验返回错误为：岁数太小，不能小于18 时", func() {
				gomega.Expect(gopher.Validate(inputData[2])).
					Should(gomega.MatchError("岁数太小，不能小于18"))
			})
		})

		Context("当获取的数据校验成功时", func() {
			It("通过了数据校验", func() {
				gomega.Expect(gopher.Validate(inputData[3])).
					Should(gomega.BeNil())
			})
		})
	})

	AfterEach(func() {
		By("当测试不通过时，我会在这里打印一个消息 【AfterEach】")
	})
})
