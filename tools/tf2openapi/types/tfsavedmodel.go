package types

/**
TFSavedModel is the high level serialization format for TensorFlow saved models.
It is the internal model representation for the SavedModel defined in the TensorFlow repository
[tensorflow/core/protobuf/saved_model.proto]
*/
import (
	"errors"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/kubeflow/kfserving/pkg/utils"
	pb "github.com/kubeflow/kfserving/tools/tf2openapi/generated/protobuf"
)

const ServingMetaGraphTag string = "serve"

type TFSavedModel struct {
	MetaGraph TFMetaGraph
}

func NewTFSavedModel(model *pb.SavedModel, sigDefKey string) (TFSavedModel, error) {
	for _, metaGraph := range model.MetaGraphs {
		if !utils.Includes(metaGraph.MetaInfoDef.Tags, ServingMetaGraphTag) {
			continue
		}
		tfMetaGraph, err := NewTFMetaGraph(metaGraph, sigDefKey)
		if err != nil {
			return TFSavedModel{}, err
		}
		return TFSavedModel{
			MetaGraph: tfMetaGraph,
		}, nil
	}
	return TFSavedModel{}, errors.New("model does not contain any servable MetaGraphs")
}

func (t *TFSavedModel) Schema() *openapi3.Schema {
	return t.MetaGraphs[0].Schema()
}
