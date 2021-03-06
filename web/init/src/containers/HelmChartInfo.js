import { connect } from "react-redux";
import realHelmChartInfo from "../components/kustomize/HelmChartInfo";

import { submitAction } from "../redux/data/determineSteps/actions";
import { getMetadata } from "../redux/data/kustomizeSettings/actions";

const HelmChartInfo = connect(
  state => ({
    dataLoading: state.ui.main.loading,
    shipAppMetadata: state.data.kustomizeSettings.shipAppMetadata,
    actions: state.data.determineSteps.stepsData.actions,
  }),
  dispatch => ({
    getMetadata() { return dispatch(getMetadata()) },
    submitAction(action) { return dispatch(submitAction(action)); }
  }),
)(realHelmChartInfo);

export default HelmChartInfo;