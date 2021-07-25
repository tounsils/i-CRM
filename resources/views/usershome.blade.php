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

        <div class="container">

        @if (session('successMsg'))
        <div class="alert alert-success" role="alert">
            {{ session('SuccessMsg') }}
        </div>
        @endif

        <div class="row">


    <table class="table table-striped">
    <thead>
        <tr>
          <th>ID</th>
          <th>First</th>
          <th>Last</th>
          <th>Email</th>
          <th>Phone</th>
          <th colspan = 2>Actions</th>
        </tr>
    </thead>
    <tbody>
    @foreach($users as $user)
        <tr>
          <td>{{ $user->id }}</td>
          <td>{{ $user->first_name }}</td>
          <td>{{ $user->last_name }}</td>
          <td>{{ $user->email }}</td>
          <td>{{ $user->phone }}</td>
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