@extends('layouts.main')

@section('title', 'Users Home')

@section('content')
 
<main id="main">
<hr>

    <!-- ======= user Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Users</h2>
          <p>Listing</p>
        </header>

        @if (session('successMsg'))
        <div class="alert alert-success" role="alert">
            {{ session('SuccessMsg') }}
        </div>
        @endif

        <div class="row">
<div class="col-sm-12">
    <h1 class="display-3">Contacts</h1>    
  <table class="table table-striped">
    <thead>
        <tr>
        <td >#</td>
        <td >First</td>
        <td >Last</td>
        <td >Email</td>
        <td >Phone</td>
        <td colspan = 2>Actions</td>
        </tr>
    </thead>
    <tbody>
    @foreach($users as $user)
        <tr>
        <td>{{ $user->id }}<td>
        <td>{{ $user->first_name }}<tdd>
        <td>{{ $user->last_name }}<td>
        <td>{{ $user->email }}<td>
        <td>{{ $user->phone }}<td>
        <td>
                <a href="{{ route('users.edit',$user->id)}}" class="btn btn-edit">Edit</a>
            </td>
        <td>
                <form action="{{ route('users.destroy', $user->id)}}" method="post">
                  @csrf
                  @method('DELETE')
                  <button class="btn btn-danger" type="submit">Delete</button>
                </form>
            </td>
        </tr>
        @endforeach


    </tbody>
  </table>
<div>
</div>








      </div>

    </section><!-- End user Section -->



</main><!-- End #main -->


@endsection