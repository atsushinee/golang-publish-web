{{define "confirm-dialog"}}
<div id="dialog" class="modal fade" aria-hidden="true" tabindex="-1"
     role="dialog"
     aria-labelledby="myModalLabel" style="margin-top: 120pt;">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header text-center">
                <h4 class="modal-title text-primary"><strong id="dialog-title">提示</strong></h4>
            </div>
            <div class="modal-body text-center">
                <div class="form-group form-inline">
                    <span class="text-danger" id="dialog-content"></span>
                </div>

            </div>
            <div class="modal-footer">
                <div class="text-center">
                    <button class="btn btn-danger" data-dismiss="modal">返回</button>
                    <button class="btn btn-success" id="dialog-submit" style="margin-left: 30pt;">确定
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>
{{end}}