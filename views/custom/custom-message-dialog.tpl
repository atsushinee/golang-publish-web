{{define "message-dialog"}}
<div id="dialog" class="modal fade" aria-hidden="true" tabindex="-1"
     role="dialog"
     aria-labelledby="myModalLabel" style="margin-top: 120pt;">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header text-center">
                <h4 class="modal-title text-primary"><strong id="dialog-title">提示</strong></h4>
            </div>
            <div class="modal-body text-center" style="margin: 2em">
                <span class="text-danger" id="dialog-content"></span>
            </div>
        </div>
    </div>
</div>
{{end}}