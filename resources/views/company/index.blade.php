@extends('layouts.main')

@section('title', 'Company Home')

@section('content')
 
<main id="main">

    <!-- ======= company Section ======= -->
    <section id="company" class="company">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Company</h2>
          <p>Listing</p>
        </header>

<div class="message" data-aos="fade-up">
@if (session('message'))
<div class="alert session('alert-type')">
{!! session('message') !!}
</div>
@endif
</div>

    <a class="btn btn-success" href="javascript:void(0)" id="createNewcompany"> 
    <i class="ri-add-line"></i>  Create New company</a>
    <table class="table table-striped">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Email</th>
                <th>Phone</th>
                <th colspan = 2>Actions</th>
            </tr>
        </thead>
        <tbody>

        @foreach($companies as $company)
        <tr>
            <td>{{$company->id}}</td>
            <td>{{$company->company_name}} </td>
            <td>{{$company->email}}</td>
            <td>{{$company->phone}}</td>
            <td>
                <a href="{{ route('company.edit',$company->id)}}"
                class="btn btn-warning btn-edit">Edit  <i class="ri-pencil-line"></i> </a>
            </td>
            <td>
                <form action="{{ route('company.destroy', $company->id)}}" 
                    id="delete-form-{{ $company->id }}"
                    method="post">
                  @csrf
                  @method('DELETE')
                  <button onclick="if (confirm('Are you sure to delete this data?')) {
                    event.preventDefault();
                    document.getElementById('delete-form-{{ $company->id }}').submit();
                  }else{
                      event.preventDefault();
                  }
                    " 
                  class="btn btn-danger" type="submit">Delete <i class="ri-delete-bin-line"></i>  </button>
                </form>
            </td>
        </tr>
        @endforeach
    </tbody>
    </table>
    {{-- Pagination --}}
        <div class="d-flex justify-content-center">
            {!! $companies->links() !!}
        </div>

</div>

<div class="modal fade" id="ajaxModel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title" id="modelHeading"></h4>
            </div>
            <div class="modal-body">
                <form id="companyForm" name="companyForm" class="form-horizontal" 
                action="{{ route('company.store') }}" method="post">
                {{csrf_field()}}


                   <input type="hidden" name="company_id" id="company_id">
                   <div class="form-group">
                        <label for="company_name" class="col-sm-2 control-label">Name</label>
                        <div class="col-sm-12">
                            * <input type="text" class="form-control" id="company_name" name="company_name" placeholder="Enter Name" value="" maxlength="50" required="">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-2 control-label">Email</label>
                        <div class="col-sm-12">
                            <input type="text" class="form-control" id="email" name="email" placeholder="Enter email" value="" maxlength="50" >
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-2 control-label">Phone</label>
                        <div class="col-sm-12">
                            <input type="text" class="form-control" id="phone" name="phone" placeholder="Enter phone" value="" maxlength="50">
                        </div>
                    </div>

                    <div class="col-sm-offset-2 col-sm-10">
                     <button type="submit" class="btn btn-primary" id="saveBtn" value="create">Save changes
                     </button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>



      </div>

    </section><!-- End company Section -->



</main><!-- End #main -->


@endsection

@section('footer-scripts')
    @include('scripts.ajax')
@endsection
