package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-work-center-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-work-center-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strings"
)

func (c *DPFMAPICaller) GeneralRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	where := strings.Join([]string{
		fmt.Sprintf("WHERE general.WorkCenter = %d ", input.General.WorkCenter),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	general.WorkCenter
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_work_center_general_data as general 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
