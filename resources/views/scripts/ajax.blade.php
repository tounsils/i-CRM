  <script type="text/javascript">
  $(function () {

      $.ajaxSetup({
          headers: {
              'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
          }
    });

    var table = $('.data-table').DataTable({
        processing: true,
        serverSide: true,
        ajax: "",
        columns: [
            {data: 'DT_RowIndex', name: 'DT_RowIndex'},
            {data: 'company_name', name: 'company_name'},
            {data: 'email', name: 'email'},
            {data: 'phone', name: 'phone'},
            {data: 'action', name: 'action', orderable: false, searchable: false},
        ]
    });

    $('#createNewcompany').click(function () {
        $('#saveBtn').val("create-company");
        $('#company_id').val('');
        $('#companyForm').trigger("reset");
        $('#modelHeading').html("Create New company");
        $('#ajaxModel').modal('show');
    });

    $('body').on('click', '.editcompany', function () {
      var company_id = $(this).data('id');
      $.get("" +'/' + company_id +'/edit', function (data) {
          $('#modelHeading').html("Edit company");
          $('#saveBtn').val("edit-user");
          $('#ajaxModel').modal('show');
          $('#company_id').val(data.id);
          $('#name').val(data.name);
          $('#detail').val(data.detail);
      })
   });

    $('#saveBtn').click(function (e) {
        e.preventDefault();
        $(this).html('Sending..');

        $.ajax({
          data: $('#companyForm').serialize(),
          url: "",
          type: "POST",
          dataType: 'json',
          success: function (data) {

              $('#companyForm').trigger("reset");
              $('#ajaxModel').modal('hide');
              table.draw();

          },
          error: function (data) {
              console.log('Error:', data);
              $('#saveBtn').html('Save Changes');
          }
      });
    });

    $('body').on('click', '.deletecompany', function () {

        var company_id = $(this).data("id");
        confirm("Are You sure want to delete !");

        $.ajax({
            type: "DELETE",
            url: ""+'/'+company_id,
            success: function (data) {
                table.draw();
            },
            error: function (data) {
                console.log('Error:', data);
            }
        });
    });

  });
</script>
